import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";
import { ReactComponent as Logo } from "./smaillogo.svg";

const HeaderWrappper = styled.div`
  padding-top: 40px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  font-size: 1em;
  text-align: center;
`;

const HeaderItem = styled.div`
  padding: 5px;
  text-decoration: none;
`;

const Header = () => {
  return (
    <HeaderWrappper>
      <HeaderItem>
        <Logo />
      </HeaderItem>
      <HeaderItem>
        <Link style={{ color: "inherit", textDecoration: "inherit" }} to="/">
          Home
        </Link>
      </HeaderItem>
      <HeaderItem>
        <Link
          style={{ color: "inherit", textDecoration: "inherit" }}
          to="/unsub"
        >
          Subscriptions
        </Link>
      </HeaderItem>
      <HeaderItem>
        <Link
          style={{ color: "inherit", textDecoration: "inherit" }}
          to="/auth"
        >
          Manage your Account
        </Link>
      </HeaderItem>
    </HeaderWrappper>
  );
};

export default Header;
