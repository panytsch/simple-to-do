import React from 'react';
import { Provider } from "react-redux";
import { BrowserRouter as Router, Route } from "react-router-dom";

import store from "./redux/store";
import {StartPage} from "./components/StartPage";
import './App.css';

const App: React.FC = () => {
  return (
      <Provider store={store}>
          <Router>
              <div>
                  <Route exact path="/" component={StartPage} />
              </div>
          </Router>
      </Provider>
  );
};

export default App;
