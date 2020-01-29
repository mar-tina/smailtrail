import React, { useState, useEffect } from "react";
import axios from "axios";
import styled from "styled-components";

const Wrapper = styled.section`
  padding: 2em;
  display: grid;
  justify-content: center;
  justify-items: center;
`;

const GoogleLink = styled.a`
  text-decoration: none;
  padding: 10px;
  color: blue;
`;

const InputWrapper = styled.input`
  padding: 20px;
  margin-top: 10px;
  margin-bottom: 10px;
`;

const SubmitButton = styled.button`
  padding: 20px;
  background: red;
  color: white;
  font-size: 15px;
  margin-top: 20px;
`;

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
        <div>
          <p> Authenticate with google and then paste the given code below </p>

          <GoogleLink href={authURL} target="_blank">
            Google Auth
          </GoogleLink>
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

          <div>
            {authresponse !== "" ? <div> {authresponse} </div> : <div> </div>}
          </div>
        </div>
      )}
    </Wrapper>
  );
};

export default Auth;
