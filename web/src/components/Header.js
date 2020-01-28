import React from "react";
import styled from "styled-components";
import { Link } from "react-router-dom";

const HeaderWrappper = styled.div`
  padding: 20px;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
`;

const HeaderItem = styled.div`
  padding: 5px;
  text-decoration: none;
`;

const Header = () => {
  return (
    <HeaderWrappper>
      <HeaderItem>
        <Link to="/">Home</Link>
      </HeaderItem>
      <HeaderItem>
        <Link to="/unsub">Subscriptions</Link>
      </HeaderItem>
      <HeaderItem>
        <Link to="/auth">Auth</Link>
      </HeaderItem>
    </HeaderWrappper>
  );
};

export default Header;
