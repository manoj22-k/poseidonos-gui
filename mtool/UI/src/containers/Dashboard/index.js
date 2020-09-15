/* -------------------------------------------------------------------------------------/
                                                                                    /
/               COPYRIGHT (c) 2019 SAMSUNG ELECTRONICS CO., LTD.                      /
/                          ALL RIGHTS RESERVED                                        /
/                                                                                     /
/   Permission is hereby granted to licensees of Samsung Electronics Co., Ltd.        /
/   products to use or abstract this computer program for the sole purpose of         /
/   implementing a product based on Samsung Electronics Co., Ltd. products.           /
/   No other rights to reproduce, use, or disseminate this computer program,          /
/   whether in part or in whole, are granted.                                         / 
/                                                                                     /
/   Samsung Electronics Co., Ltd. makes no representation or warranties with          /
/   respect to the performance of this computer program, and specifically disclaims   /
/   any responsibility for any damages, special or consequential, connected           /
/   with the use of this program.                                                     /
/                                                                                     /
/-------------------------------------------------------------------------------------/


DESCRIPTION: <File description> *
@NAME : index.js
@AUTHORS: Jay Hitesh Sanghavi 
@Version : 1.0 *
@REVISION HISTORY
[03/06/2019] [Jay] : Prototyping..........////////////////////
*/
import React, { Component } from 'react';
import { connect } from 'react-redux'
import MaterialTable from 'material-table';
import 'react-dropdown/style.css';
import { withStyles } from '@material-ui/core/styles';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import FirstPage from '@material-ui/icons/FirstPage';
import LastPage from '@material-ui/icons/LastPage';
import Remove from '@material-ui/icons/Remove';
import ArrowUpward from '@material-ui/icons/ArrowUpward';
import 'react-table/react-table.css';
import 'core-js/es/number';
import 'core-js/es/array';
import { Paper, Grid, Typography } from '@material-ui/core';
import ThemeProvider from '@material-ui/core/styles/MuiThemeProvider';
import formatBytes from '../../utils/format-bytes';
import { customTheme, PageTheme } from '../../theme';
import Header from '../../components/Header';
import Sidebar from '../../components/Sidebar';
import './Dashboard.css';
import * as actionTypes from "../../store/actions/actionTypes";
import * as actionCreators from "../../store/actions/exportActionCreators";
import Legend from '../../components/Legend';
import bytesToTB from '../../utils/bytes-to-tb';

const styles = (theme) => {
  return ({
  dashboardContainer: {
    display: 'flex',
  },
  storageDetailsPaper: {
    height: 250,
    position: "relative",
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(1),
    paddingLeft:"35px",
    paddingRight:"35px",
    width: '100%',
    boxSizing: 'border-box'
  },
  toolbar: customTheme.toolbar,
  titleContainer: {
    marginTop: theme.spacing(1)
  },
  tableContainer: {
    minHeight: '372px'
  },
  loadWrapper: {
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center'
  },
  metricContainer: {
    marginBottom: theme.spacing(1)
  },
  volumeSummary: {
    minHeight: 413,
    [theme.breakpoints.down('md')]: {
      minHeight: 460
    }
  },
  metricBox: {
    display: 'flex',
    alignItems: 'center',
    width: '100%',
    height: '110px',
    justifyContent: 'center',
    borderRadius: '4px'
  },
  storageGraph: {
    position: 'absolute',
    height: '100%',
    top: 0
  },
  metricTxt: {
    color: '#fff'
  },
  spaced: {
    marginTop: theme.spacing(1)
  },
  topGrid:{
    marginBottom: '-8px'
  },
  cardHeader: {
    ...customTheme.card.header,
    marginLeft: 0
  },
  pageHeader: {
    textAlign: 'left',
    fontSize: '16px',
    fontWeight: 'bold',
    // color: 'rgb(53, 85, 142)',
    color: '#424850',
  },
  textOverflow: {
    whiteSpace: 'nowrap',
    textOverflow: 'ellipsis',
    overflow: 'hidden'
  },
  metricsPaper: {
    minHeight: 203,
    display: 'flex',
    alignItems: 'center'
  }
})
};

/* eslint-disable react/no-multi-comp */
const icons = {
  FirstPage: () => <FirstPage id="Dashboard-icon-vol-firstpage" /> ,
  LastPage: () => <LastPage id="Dashboard-icon-vol-lastpage" />,
  NextPage: () => <ChevronRight id="Dashboard-icon-vol-nextpage" />,
  PreviousPage: () => <ChevronLeft id="Dashboard-icon-vol-previouspage" />,
  ThirdStateCheck: Remove,
  DetailPanel: ChevronRight,
  SortArrow: ArrowUpward
};

/* eslint-enable react/no-multi-comp */

// Disabling for PoC1
// const alertIcons = {
//   ...icons,
//   FirstPage: () => <FirstPage id="Dashboard-alert-vol-firstpage" /> ,
//   LastPage: () => <LastPage id="Dashboard-alert-vol-lastpage" />,
//   NextPage: () => <ChevronRight id="Dashboard-alert-vol-nextpage" />,
//   PreviousPage: () => <ChevronLeft id="Dashboard-alert-vol-previouspage" />,
// };


// eslint-disable-next-line react/no-multi-comp
class Dashboard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      time: '',
      mobileOpen: false
    };
    this.interval = null;
    this.handleDrawerToggle = this.handleDrawerToggle.bind(this);
  }

  componentDidMount() {
    this.props.fetchVolumes();
    this.props.fetchStorageInfo();
    this.props.fetchPerformance();
    this.props.fetchIpAndMacInfo();
    this.props.fetchAlertsInfo();
    const today = new Date();
    const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    const date = `${months[today.getMonth()]} ${today.getDate()}, ${today.getFullYear()},`;
    const time = `${today.getHours()}:${today.getMinutes()}:${today.getSeconds()}`;
    const dateTime = `${date} ${time}`;
    this.setState({ time: dateTime });
    this.props.enableFetchingAlerts(true);
    this.interval = setInterval(() => {
      this.props.fetchPerformance();
    }, 2000);
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }


  handleDrawerToggle() {
    this.setState({
      mobileOpen: !this.state.mobileOpen
    });
  }

  render() {
    let volUsedSpace = 0; let volSpace = 0;
    this.props.volumes.forEach((vol) => {
      volUsedSpace += Number(vol.total) - Number(vol.remain);
      volSpace += Number(vol.total);
    });
    const volFilledStyle = {
      width: `${(volSpace * 100/this.props.arraySize)}%`,
      height: '100%',
      backgroundColor: '#e0e0e0',
      float: 'left',
    };
    const volUsedStyle = {
      width: `${(volUsedSpace * 100/volSpace)}%`,
      height: '100%',
      backgroundColor: 'rgba(0, 186, 0,0.6)',
      float: 'left',
    };
    const storageFreeStyle = {
      width: `${100 - (volSpace * 100/this.props.arraySize)}%`,
      height: '100%',
      color: 'black',
      marginLeft: '0px',
      float: 'left',
      overflowY: 'auto',
      display: 'inline-block',
      textAlign: 'center',
      position: 'relative',
      backgroundColor: '#fff',
    };
    const storageDangerStyle = {
      width: '10%',
      right: '0px',
      height: '100%',
      float: 'right',
      position: 'absolute',
      display: 'inline-block',
      borderLeft: '2px solid rgb(255, 173, 173)'
    };

    // Disabling for PoC1
    // const alertColumns = [
    //   {
    //     title: 'Alert Name',
    //     field: 'alertName',
    //     sorting: false,
    //   },
    //   {
    //     title: 'Time Stamp',
    //     field: 'time',
    //     defaultSort: 'desc',
    //     sorting: false,
    //   },
    //   {
    //     title: 'Status',
    //     field: 'level',
    //     width: 100,
    //     sorting: false,
    //     render: row => {
    //       if (row.level === 'CRITICAL') {
    //         return <span style={{ color: 'red' }}>{row.level}</span>;
    //       }
    //       return <span style={{ color: 'green' }}>{row.level}</span>;
    //     },
    //   },
    //   {
    //     title: 'Description',
    //     field: 'message',
    //     sorting: false,
    //   },
    //   {
    //     title: 'Duration(sec)',
    //     field: 'duration',
    //     width: 150,
    //     sorting: false,
    //   },
    //  // {
    //   //  title: 'Node',
    //    // field: 'host',
    //     // sorting: false,
    //  // },
    // ];
    const volumeTableColumns = [
      {
        title: 'Name',
        field: 'name',
      },
      {
        title: 'Used Space (GB)',
	render: rowData => rowData.usedspace ?
            formatBytes(rowData.total - rowData.remain) : 0
      },
      {
        title: 'Total Space',
        render: rowData => rowData.total ?
            formatBytes(rowData.total) : 0
      },
    ];
    const { classes } = this.props;
    return (
      <ThemeProvider theme={PageTheme}>
      <div className={classes.dashboardContainer}>
        <Header toggleDrawer={this.handleDrawerToggle} />
        <Sidebar mobileOpen={this.state.mobileOpen} toggleDrawer={this.handleDrawerToggle} />
        <main className={classes.content}>  
          <div className={classes.toolbar} />
          <Grid container spacing={3}>
          <Grid container spacing={3} className={classes.titleContainer}>
            <Grid sm={6} xs={12} item>
              <Typography className={classes.pageHeader} variant="h6">Dashboard</Typography>
            </Grid>
            <Grid md={6} xs={12} item container direction="row" justify="space-between">
              <Typography data-testid="dashboard-ip" variant="body1" component="span">
                  IP : {this.props.ip}
              </Typography>
              <Typography data-testid="dashboard-host" variant="body1" component="span">
                  Poseidon Name : {this.props.host}
              </Typography>
              <Typography data-testid="dashboard-mac" variant="body1" component="span">
                  MAC : {this.props.mac}
              </Typography>
            </Grid>
          </Grid>
          <Grid container spacing={1} className={classes.topGrid}>
            <Grid xs={12} md={6} item>
              <Paper spacing={3} className={classes.metricsPaper}>
                <Grid container justify="space-around">
                     <Grid xs={10} md={3} className={classes.metricContainer} item spacing-xs-1="true">
                       <Typography align="center" className={classes.textOverflow} color="secondary">BANDWIDTH</Typography>
                       <Grid item xs={12} className={classes.metricBox} style={{backgroundColor: 'rgba(58, 108, 255,0.6'}}>
                        <Typography variant="h2" data-testid="dashboard-bandwidth" className={classes.metricTxt}>{this.props.bw} MBps</Typography>
                       </Grid>
                     </Grid>
                     <Grid xs={10} md={3} className={classes.metricContainer} item spacing-xs-1="true">
                       <Typography align="center" className={classes.textOverflow} color="secondary">READ IOPS</Typography>
                       <Grid item xs={12} className={classes.metricBox} style={{backgroundColor: 'rgba(59, 189, 179,0.7'}}>
                       <Typography 
                       variant="h2" 
                       data-testid= "read-iops"
                       className={classes.metricTxt}
                       >
                       {this.props.read_iops}
                       </Typography>
                       </Grid>
                     </Grid>
                     <Grid xs={10} md={3} className={classes.metricContainer} item spacing-xs-1="true">
                       <Typography align="center" className={classes.textOverflow} color="secondary">WRITE IOPS</Typography>
                       <Grid item xs={12} className={classes.metricBox} style={{backgroundColor: 'rgba(228, 148, 42,0.6'}}>
                       <Typography variant="h2" data-testid= "write-iops" className={classes.metricTxt}>{this.props.write_iops}</Typography>
                       </Grid>
                     </Grid>
                </Grid>
              </Paper>
              <Paper spacing={3} className={`${classes.spaced} ${classes.storageDetailsPaper}`}>
                <Grid container justify="space-between">
                  <Typography className={classes.cardHeader} style={{marginLeft: "24px"}}>
                    Storage Details
                  </Typography>
                </Grid>
                <Grid container justify="center" alignContent="center" className={classes.storageGraph}>
                {this.props.arraySize === 0 ? (
                  <Typography data-testid="dashboard-no-array"color="secondary">No Array Created</Typography>
                ) : (
                  <React.Fragment>
                    <div className="dashboard-size-label-container">
                      <span className="dashboard-min-label">0TB</span>
                      <span className="dashboard-max-label">
                        {bytesToTB(this.props.arraySize)}
                      </span>
                    </div>
                    <div className="storage-detail-container">
                      <div style={volFilledStyle}>
                        <div style={volUsedStyle} />
                      </div>
                      <div style={storageFreeStyle} />
                      <div style={storageDangerStyle}>
                          <div className="dashboard-threshold-label">80%</div>
                      </div>
                    </div>
                    <div
                      style={{
                        width: '94%',
                        margin: '5px auto auto',
                        height: 'auto',
                        position: 'relative',
                      }}
                    >
                      <Grid item container xs={12} wrap="wrap">
                      <Legend
                        bgColor="rgba(0, 186, 0, 0.6)"
                        title={`Data Written: ${bytesToTB(volUsedSpace)}`}
                      />
                      <Legend
                        bgColor="#e0e0e0"
                        title={`Volume Space Allocated: ${bytesToTB(volSpace)}`}
                      />
                      <Legend
                        bgColor="#fff"
                        title={`Available for Volume Creation: ${(bytesToTB(this.props.arraySize - volSpace))}`}
                      />
                      <Legend
                        bgColor="rgb(255, 173, 173)"
                        title="Threshold"
                      />
                      </Grid>
                      <span
                        style={{
                          width: '100%',
                          marginTop: '10px',
                          float: 'left',
                          textAlign: 'left',
                        }}
                      >
                        As of {this.state.time}
                      </span>
                    </div>
                  </React.Fragment>
                )}
                </Grid>
              </Paper>
            </Grid>

            <Grid xs={12} md={6} item>
              <Paper spacing={3} className={classes.volumeSummary}>
                <MaterialTable
                  title={(
                    <Typography className={classes.cardHeader}>Volume Summary</Typography>
                  )}
                  columns={volumeTableColumns}
                  data={this.props.volumes}
                  options={{
                    headerStyle: {
                      backgroundColor: '#788595',
                      color: '#FFF'
                    },
                    minBodyHeight: 342,
                    maxBodyHeight: 342,
                    search: false
                  }}
                  style={{height: '100%'}}
                  icons={icons}
                />
              </Paper>
            </Grid>
          </Grid>
          {/* <Grid container spacing={1} className={classes.spaced}>
            <Grid xs={12} item>
            <Paper>
              {this.props.alerts.length > 0 ||
                this.props.fetchingAlerts === false ? (
                <MaterialTable
                  title={(
                    <Typography className={classes.cardHeader}>Storage Alerts</Typography>
                  )}
                  data-testid="Dashboard-table-alert"
                  options={{
                    headerStyle: {
                      backgroundColor: '#788595',
                      color: '#FFF'
                    },
                    search: false
                  }}
                  data={this.props.alerts}
                  columns={alertColumns}
                  icons={alertIcons}
                />
              ) : (
                <div className={classes.loadWrapper}>
                  <Loader
                    type="Bars"
                    color="#788595"
                    marginTop="100px"
                    width="50"
                  />
                </div>
              )}
            </Paper>
            </Grid>
          </Grid> */}
          </Grid>
        </main>
      </div>
      </ThemeProvider>
    );
  }
}
const mapStateToProps = state => {
  return {
    volumes: state.dashboardReducer.volumes,
    alerts: state.dashboardReducer.alerts,
    ibofs: state.dashboardReducer.ibofs,
    unusedSpace: state.dashboardReducer.unusedSpace,
    used: state.dashboardReducer.used,
    unused: state.dashboardReducer.unused,
    read_iops: state.dashboardReducer.read_iops,
    write_iops: state.dashboardReducer.write_iops,
    bw: state.dashboardReducer.bw,
    fetchingAlerts: state.dashboardReducer.fetchingAlerts,
    ip: state.dashboardReducer.ip,
    mac: state.dashboardReducer.mac,
    host: state.dashboardReducer.host,
    arraySize: state.dashboardReducer.arraySize
  };
}
const mapDispatchToProps = dispatch => {
  return {
    enableFetchingAlerts: (flag) => dispatch(actionCreators.enableFetchingAlerts(flag)),
    fetchVolumes: () => dispatch({ type: actionTypes.SAGA_FETCH_VOLUME_INFO }),
    fetchAlertsInfo: () => dispatch({ type: actionTypes.SAGA_FETCH_ALERTS_INFO }),
    fetchStorageInfo: () => dispatch({ type: actionTypes.SAGA_FETCH_STORAGE_INFO }),
    fetchPerformance: () => dispatch({ type: actionTypes.SAGA_FETCH_PERFORMANCE_INFO }),
    fetchIpAndMacInfo: () => dispatch({ type: actionTypes.SAGA_FETCH_IPANDMAC_INFO }),
  };
}
export default withStyles(styles)((connect(mapStateToProps, mapDispatchToProps))(((Dashboard))));
