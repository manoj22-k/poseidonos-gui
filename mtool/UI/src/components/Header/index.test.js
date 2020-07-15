/* eslint-disable import/imports-first */
/* eslint-disable import/first */

jest.unmock('axios');

import axios from "axios";
import MockAdapter from "axios-mock-adapter";
import React from 'react';
import { Provider } from 'react-redux';
import { Router } from 'react-router-dom';
import { render, fireEvent, cleanup, waitForElement, getByPlaceholderText, getByText } from '@testing-library/react';
import { createMemoryHistory } from 'history';
import '@testing-library/jest-dom/extend-expect';
import { createStore, combineReducers, applyMiddleware, compose } from 'redux';
import createSagaMiddleware from 'redux-saga';
import { act } from "react-dom/test-utils";
import Header from './index';
import headerReducer from '../../store/reducers/headerReducer';
import configurationsettingReducer from '../../store/reducers/configurationsettingReducer';
import rootSaga from "../../sagas/indexSaga";
import { async } from "q";


describe('<Header />', () => {
  let wrapper;
  let mock;
  let history;
  let store;
  beforeEach(() => {
    const sagaMiddleware = createSagaMiddleware();
    const rootReducers = combineReducers({
      headerReducer,
      configurationsettingReducer
    });
    const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
    store = createStore(rootReducers, composeEnhancers(applyMiddleware(sagaMiddleware)))
    sagaMiddleware.run(rootSaga);
    const route = '/';
    history = createMemoryHistory({ initialEntries: [route] })
    mock = new MockAdapter(axios);
  });

  const renderComponent = () => {
    wrapper = render(
      <Router history={history}>
        <Provider store={store}>
          <Header />
        </Provider>
      </Router>
    );
  }

  afterEach(cleanup);

  it("should render Change Password dialogue", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    const oldPwd = await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    expect(oldPwd).toBeDefined();
    // expect(asFragment()).toMatchSnapshot();
    fireEvent.click(await waitForElement(() => getByText('Poseidon OS status:')));
  });

  it("should change the Password", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    const oldPwd = await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    fireEvent.keyDown(oldPwd, { key: 'A', code: 65, charCode: 65 });
    fireEvent.change(oldPwd, {target: {value: "abcd"}});
    const newPwd = await waitForElement(() => getByPlaceholderText("Enter New Password"));
    fireEvent.keyDown(newPwd, { key: 'D', code: 68, charCode: 68 });
    fireEvent.change(newPwd, {target: {value: "defg"}});
    const confPwd = await waitForElement(() => getByPlaceholderText("Confirm New Password"));
    fireEvent.keyDown(confPwd, { key: 'D', code: 68, charCode: 68 });
    fireEvent.change(confPwd, {target: {value: "defg"}});
    jest.spyOn(global, 'fetch')
    .mockImplementation(() => Promise.resolve({
      status: 200,
      json: () => Promise.resolve({
        value: ""
      })
    }));
    fireEvent.click(await waitForElement(() => getByTestId("change-pwd-submit")));
  });

  it("should redirect if user session ended", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    const oldPwd = await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    fireEvent.keyDown(oldPwd, { key: 'A', code: 65, charCode: 65 });
    fireEvent.change(oldPwd, {target: {value: "abcd"}});
    const newPwd = await waitForElement(() => getByPlaceholderText("Enter New Password"));
    fireEvent.keyDown(newPwd, { key: 'D', code: 68, charCode: 68 });
    fireEvent.change(newPwd, {target: {value: "defg"}});
    const confPwd = await waitForElement(() => getByPlaceholderText("Confirm New Password"));
    fireEvent.keyDown(confPwd, { key: 'D', code: 68, charCode: 68 });
    fireEvent.change(confPwd, {target: {value: "defg"}});
    jest.spyOn(global, 'fetch')
    .mockImplementation(() => Promise.resolve({
      status: 401,
      json: () => Promise.resolve({
        value: ""
      })
    }));
    fireEvent.click(await waitForElement(() => getByTestId("change-pwd-submit")));
  });

  it("should throw error if password change failed", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    const oldPwd = await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    fireEvent.keyDown(oldPwd, { key: 'A', code: 65, charCode: 65 });
    fireEvent.change(oldPwd, {target: {value: "abcd"}});
    const newPwd = await waitForElement(() => getByPlaceholderText("Enter New Password"));
    fireEvent.keyDown(newPwd, { key: 'D', code: 68, charCode: 68 });
    fireEvent.change(newPwd, {target: {value: "defg"}});
    const confPwd = await waitForElement(() => getByPlaceholderText("Confirm New Password"));
    fireEvent.keyDown(confPwd, { key: 'D', code: 68, charCode: 68 });
    fireEvent.change(confPwd, {target: {value: "defg"}});
    jest.spyOn(global, 'fetch')
    .mockImplementation(() => Promise.resolve({
      status: 400,
      json: () => Promise.resolve({
        value: ""
      })
    }));
    fireEvent.click(await waitForElement(() => getByTestId("change-pwd-submit")));
    expect(await waitForElement(() => getByText("Error in setting Password"))).toBeDefined();
  });

  it("should show error when old password is not entered", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    fireEvent.click(await waitForElement(() => getByTestId("change-pwd-submit")));
    fireEvent.click(await waitForElement(() => getByText('OK')));
  });

  it("should show error when new password is not entered", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    const oldPwd = await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    fireEvent.change(oldPwd, {target: {value: "abcd"}});
    fireEvent.click(await waitForElement(() => getByTestId("change-pwd-submit")));
    fireEvent.click(await waitForElement(() => getByText('OK')));
  });

  it("should show error when confirm password is not entered", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    const oldPwd = await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    fireEvent.change(oldPwd, {target: {value: "abcd"}});
    const newPwd = await waitForElement(() => getByPlaceholderText("Enter New Password"));
    fireEvent.change(newPwd, {target: {value: "defg"}});
    fireEvent.click(await waitForElement(() => getByTestId("change-pwd-submit")));
    fireEvent.click(await waitForElement(() => getByText('OK')));
  });

  it("should show error when new password and confirm password do not match", async () => {
    renderComponent();
    const  { asFragment, getByTestId, getByText, getByPlaceholderText } = wrapper;
    fireEvent.click(getByTestId('header-dropdown'));
    fireEvent.click(await waitForElement(() => getByText('Change Password')));
    const oldPwd = await waitForElement(() => getByPlaceholderText("Enter Old Password"));
    fireEvent.change(oldPwd, {target: {value: "abcd"}});
    const newPwd = await waitForElement(() => getByPlaceholderText("Enter New Password"));
    fireEvent.change(newPwd, {target: {value: "deg"}});
    const confPwd = await waitForElement(() => getByPlaceholderText("Confirm New Password"));
    fireEvent.change(confPwd, {target: {value: "defg"}});
    fireEvent.click(await waitForElement(() => getByTestId("change-pwd-submit")));
    expect(await waitForElement(() => getByText("Passwords do not match"))).toBeDefined();
  });

  it("should toggle the menu", async () => {
    window.innerWidth = 500
    window.innerHeight = 200
    window.dispatchEvent(new Event('resize'));
    renderComponent();
    const { getByText, asFragment, getByTestId } = wrapper;
    fireEvent.click(await waitForElement(() => getByTestId('sidebar-toggle')));
    // expect(await waitForElement(() => getByText('Dashboard'))).toBeDefined();
    fireEvent.click(await waitForElement(() => getByTestId('mobile-show-more')));
    fireEvent.click(await waitForElement(() => getByTestId('menu-expand')));
    fireEvent.click(await waitForElement(() => getByText('Poseidon OS status:')));
    expect(asFragment()).toMatchSnapshot();
    // const dashboardLink = expect(await waitForElement(() => getByText('Poseidon OS status:')));
    // fireEvent.click(dashboardLink);
  });

//   it("should adjust the dropdown position", async () => {
//     renderComponent();
//     const { getByText, asFragment } = wrapper;
//     window.innerWidth = 100
//     window.innerHeight = 200
//     window.dispatchEvent(new Event('resize'));
//     expect(asFragment()).toMatchSnapshot();
//     const dashboardLink = expect(await waitForElement(() => getByText('Poseidon OS status:')));
//     fireEvent.click(dashboardLink);
//   });
});