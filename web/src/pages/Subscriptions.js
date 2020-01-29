import React, { useState, useEffect } from "react";
import styled from "styled-components";
import axios from "axios";

const Title = styled.div`
  font-size: 1.5em;
  text-align: center;
  color: palevioletred;
`;

const Wrapper = styled.section`
  padding: 2em;
  display: grid;
  justify-content: center;
`;

const FetchMoreButton = styled.button`
  padding: 10px;
  font-size: 15px;
  color: white;
  background: lightcoral;
  border: none;
`;

const MiniHeader = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-content: center;
  margin: 20px;
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
    fetchFromDB();
    fetchData();
  }, []);

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

  const fetchFromDB = async () => {
    setLoading(true);
    const result = await axios(`http://localhost:8000/subs?dbkey=${dbkey}`);

    const data = result.data;
    setdbData(dbdata => [...dbdata, data]);
    setLoading(false);
  };

  const handleFetchMore = () => {
    fetchData();
  };

  console.log("Msgs", msgs);

  return (
    <Wrapper>
      <MiniHeader>
        <Title> Manage your Subscriptions </Title>
        <div>
          <FetchMoreButton onClick={handleFetchMore}>
            Fetch More
          </FetchMoreButton>
        </div>
      </MiniHeader>

      <div>
        {dbdata.map((x, i) => (
          <div key={i}>
            {x.map((y, i) => (
              <div key={i}>
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
