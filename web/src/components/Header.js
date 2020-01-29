import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";

const HeaderWrappper = styled.div`
  padding: 20px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
  font-size: 1.5em;
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
