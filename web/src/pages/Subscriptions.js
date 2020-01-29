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

const ContentWrapper = styled.div``;

const Subscriptions = () => {
  const [msgs, setMsgs] = useState([]);
  const [npToken, setnpToken] = useState("");
  const [dbkey, setdbKey] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [from, setFrom] = useState();
  const [dbdata, setdbData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      const result = await axios(
        `http://localhost:8000/allmessages?nextpagetoken=${npToken}`
      );
      const data = result.data;
      setnpToken(data.list.nextPageToken);

      setMsgs(msgs => [...msgs, data.msgs]);
      setLoading(false);
    };
    fetchFromDB();
    fetchData();
  }, []);

  const fetchFromDB = async () => {
    setLoading(true);
    const result = await axios(`http://localhost:8000/subs?dbkey=${dbkey}`);

    const data = result.data;
    setdbData(dbdata => [...dbdata, data]);
  };

  console.log("Msgs", msgs);
  console.log("Token", npToken);
  console.log("DBData", dbdata);

  const numbers = [1, 2, 3, 4, 5];
  const subItems = dbdata.map((number, i) => (
    <div key={i}>{number.Sender}</div>
  ));

  return (
    <Wrapper>
      <Title> Manage your Subscriptions </Title>
      <div>
        {dbdata.map(x => (
          <div>
            {x.map(y => (
              <div>
                {" "}
                {y.sender}
                {y.link !== "" ? (
                  <a href={y.link}> Unsub </a>
                ) : (
                  <a href={y.link}> Missing Link </a>
                )}
              </div>
            ))}
          </div>
        ))}
      </div>
    </Wrapper>
  );
};

export default Subscriptions;
