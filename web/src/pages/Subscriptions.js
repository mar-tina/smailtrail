import React, { useState, useEffect } from "react";
import styled from "styled-components";
import axios from "axios";

const Title = styled.h1`
  font-size: 1.5em;
  text-align: center;
  color: palevioletred;
`;

const Wrapper = styled.section`
  padding: 2em;
  min-height: 500px;
`;

const ContentWrapper = styled.div`
  min-height: 500px;
`;

const Content = styled.div`
  min-height: 100px;
`;

const Subscriptions = () => {
  const [msgs, setMsgs] = useState([]);
  const [npToken, setnpToken] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [from, setFrom] = useState();

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      const result = await axios(
        `http://localhost:8000/allmessages?nextpagetoken=${npToken}`
      );
      const data = result.data;
      setMsgs(msgs => [...msgs, data.msgs]);
      setLoading(false);
    };
    fetchData();
  }, []);

  console.log("WTF", msgs);

  return (
    <Wrapper>
      {" "}
      <Title> Manage your Subscriptions </Title>
      {isLoading ? (
        <div> Loadding ... </div>
      ) : (
        <ContentWrapper>
          {" "}
          {msgs &&
            msgs.map((x, i) => (
              <div key={i}>
                {" "}
                {x &&
                  x.map((y, i) => (
                    <ContentWrapper>
                      {y.Parts &&
                        y.Parts.map((element, i) => (
                          <div key={i}>{element.body.data}</div>
                        ))}
                    </ContentWrapper>
                  ))}{" "}
              </div>
            ))}{" "}
        </ContentWrapper>
      )}
    </Wrapper>
  );
};

export default Subscriptions;
