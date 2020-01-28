import React from "react";
import Home from "./pages/Home";
import Header from "./components/Header";
import Subscriptions from "./pages/Subscriptions";
import Auth from "./pages/Auth";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

export default function App() {
  return (
    <Router>
      <div>
        <Header />

        {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/auth">
            <Auth />
          </Route>
          <Route path="/unsub">
            <Subscriptions />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}
