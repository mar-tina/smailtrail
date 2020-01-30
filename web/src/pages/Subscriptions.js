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

const SeeMoreButton = styled.button`
  padding: 10px;
  color: blue;
  background: white;
  border: none;
  font-size: 15px;
  &:hover {
    cursor: pointer;
  }
`;

const ContentWrapper = styled.div``;

const Subscriptions = () => {
  const [msgs, setMsgs] = useState([]);
  const [npToken, setnpToken] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [dbdata, setdbData] = useState([]);
  const [skip, setSkip] = useState(0);

  let take = 10;

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
    const result = await axios(
      `http://localhost:8000/subs?take=${take}&skip=${skip}`
    );

    const data = result.data;
    setdbData(dbdata => [...dbdata, data]);

    setLoading(false);
  };

  const handleFetchMore = () => {
    fetchData();
  };

  const handleDBRefetch = () => {
    let newskip = skip + take;
    setSkip(newskip);
    fetchFromDB();
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

        <SeeMoreButton onClick={handleDBRefetch}> See More </SeeMoreButton>
      </div>
    </Wrapper>
  );
};

export default Subscriptions;
