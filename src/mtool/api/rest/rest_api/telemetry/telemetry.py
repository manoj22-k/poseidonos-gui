import requests
import json
from bson import json_util
from flask import jsonify, make_response


def toJson(data):
    return json_util.dumps(data)

grafa_url = "http://localhost:3500"
ds_name = "poseidon"
PROM_AGG_QUERY_PATH = 'http://{}:{}/api/v1/query?query=sum({}{{job=\"pos\"}})'

def is_prometheusDB_running(prom_url):
    try:
        success_res = make_response("success",200)
        # Checking prometheusDB is running or not
        prom_res = requests.get(
            '{prom_url}/api/v1/status/runtimeinfo'.format(prom_url=prom_url)
        )
        prom_res = json.loads(prom_res.content)

        if "status" in prom_res and prom_res["status"] != "success":
            return make_response("prometheus DB is not running", 500)
        return success_res
            
    except Exception as e:
        return make_response("Error Occured" + repr(e), 500)


def check_telemetry_endpoint(ip, port):
    try:
        prom_url = "http://{ip}:{port}".format(ip=ip, port=port)
        response = is_prometheusDB_running(prom_url)
        if response.response[0].decode('UTF-8') != "success":
            return response
        # Checking temetry is up or not
        response = requests.get('{prom_url}/api/v1/query?query=up'.format(prom_url=prom_url))
        response = json.loads(response.content)
        if response is not None and "data" in response and "result" in response["data"] and len(response["data"]["result"]) > 0:
            for data in response["data"]["result"]:
                if "metric" in data and "job" in data["metric"] and data["metric"]["job"] == "pos" and "value" in data and len(data["value"]) == 2 and data["value"][1] == "1":
                    return jsonify({'isTelemetryEndpointUp': True})
        return jsonify({'isTelemetryEndpointUp': False})

    except Exception as e:
        return make_response("Error Occured" + repr(e), 500)


def set_telemetry_configuration(ip, port):
    try:
        success_res = make_response("success",200)
        prom_url = "http://{ip}:{port}".format(ip=ip, port=port)
        response = is_prometheusDB_running(prom_url)
        if response.response[0].decode('UTF-8') != "success":
            return response
        
        # Create data source in grafana
        url = '{grafa_url}/api/datasources'.format(grafa_url=grafa_url)
        headers = {
            "Accept": "*/*",
            "Content-Type": "application/json",
        }
        payload = {
            "name": ds_name,
            "type": "prometheus",
            "url": prom_url,
            "access": "proxy",
            "basicAuth": False
        }
        grafa_create_ds_res = requests.post(url, headers=headers, data=toJson(payload))
        grafa_create_ds_res = json.loads(grafa_create_ds_res.content)

        if "message" in grafa_create_ds_res and grafa_create_ds_res["message"] == "Datasource added":
            return success_res

        # Else Get the data source id named $ds_name
        grafa_ds_res = requests.get(url)
        grafa_ds_res = json.loads(grafa_ds_res.content)

        ds_id = -1
        if len(grafa_ds_res):
            for ds in grafa_ds_res:
                if "name" in ds and ds["name"] == ds_name and "url" in ds and ds["url"] != prom_url:
                    ds_id = ds["id"]
            
        if ds_id == -1:
            return success_res

        # Update data source with the new ip and port
        grafa_update_ds_res = requests.put(
            '{url}/{ds_id}'.format(url=url, ds_id=ds_id), 
            headers=headers, 
            data=toJson(payload)
        )
        grafa_update_ds_res = json.loads(grafa_update_ds_res.content)
        
        if "message" in grafa_update_ds_res and grafa_update_ds_res["message"] == "Datasource updated":
            return success_res
        return make_response("Unable to add data source in Grafana", 500)

    except requests.exceptions.HTTPError as errh:
        return make_response("An Http Error occurred:" + repr(errh), 503)
    except requests.exceptions.ConnectionError as errc:
        return make_response("An Error Connecting to the API occurred:" + repr(errc), 521)
    except requests.exceptions.Timeout as errt:
        return make_response("A Timeout Error occurred:" + repr(errt), 524)
    except requests.exceptions.RequestException as err:
        return make_response("An Unknown Error occurred" + repr(err), 520)
    except Exception as e:
        return make_response("Error Occured" + repr(e), 500)

def get_agg_value(ip, port, metric):
    try:
        PATH = PROM_AGG_QUERY_PATH.format(ip,port,metric)
        response = requests.get(PATH)
        response = json.loads(response.content)
        value = 0
        if response is not None and "data" in response and "result" in response["data"] and len(response["data"]["result"]) > 0 and "value" in response["data"]["result"][0] and len(response["data"]["result"][0]["value"]) == 2:
            value = int(response["data"]["result"][0]["value"][1])
        return value
    except requests.exceptions.HTTPError as http_err:
        print(f'HTTP error occurred: {http_err}')
    except Exception as err:
        print(f'Other error occurred: {err}')
