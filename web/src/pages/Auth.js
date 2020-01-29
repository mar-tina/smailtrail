import React, { useState, useEffect } from "react";
import axios from "axios";
import styled from "styled-components";

const Wrapper = styled.section`
  padding: 2em;
  display: grid;
  justify-content: center;
  justify-items: center;
  align-content: center;
  align-items: center;
`;

const GoogleLink = styled.a`
  text-decoration: none;
  padding: 10px;
  color: blue;
`;

const InputWrapper = styled.input`
  display: grid;
  justify-content: center;
  padding: 20px;
  margin: 0 auto;
  min-width: 250px;
  font-size: 15px;
`;

const SubmitButton = styled.button`
  padding: 20px;
  background: red;
  color: white;
  font-size: 15px;
`;

const BTNHolder = styled.div``;

const Auth = () => {
  const [authURL, setauthURL] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [authCode, setAuthcode] = useState("");
  const [authresponse, setAuthResponse] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      const result = await axios(`http://localhost:8000/initialauth`);
      console.log("Im here", result.data);
      setauthURL(result.data);
      setLoading(false);
    };
    fetchData();
  }, []);

  const completeAuth = async event => {
    setLoading(true);
    event.preventDefault();
    const res = await axios.post(`http://localhost:8000/completeauth`, {
      code: authCode
    });
    setAuthResponse(res.data);
    setLoading(false);
  };

  return (
    <Wrapper>
      {isLoading ? (
        <div>Loading ...</div>
      ) : (
        <>
          <p>
            {" "}
            Authenticate with{" "}
            <GoogleLink href={authURL} target="_blank">
              Google Auth
            </GoogleLink>{" "}
            and then paste the given code below{" "}
          </p>

          <br />

          <InputWrapper
            value={authCode}
            onChange={event => setAuthcode(event.target.value)}
          />

          <br />

          <SubmitButton onClick={event => completeAuth(event)}>
            {" "}
            Complete Auth{" "}
          </SubmitButton>

          <BTNHolder>
            {authresponse !== "" ? <div> {authresponse} </div> : <div> </div>}
          </BTNHolder>
        </>
      )}
    </Wrapper>
  );
};

export default Auth;
