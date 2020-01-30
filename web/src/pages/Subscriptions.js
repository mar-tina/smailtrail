import React, { useState, useEffect } from "react";
import { useRouteMatch } from "react-router-dom";
import styled from "styled-components";
import axios from "axios";
import { ReactComponent as LoadingLogo } from "./loading.svg";

const Title = styled.div`
  font-size: 1.2em;
  text-align: center;
  color: palevioletred;
`;

const Wrapper = styled.section`
  padding: 0.5em;
  display: grid;
  justify-content: center;
`;

const FetchMoreButton = styled.button`
  padding: 15px;
  font-size: 15px;
  color: white;
  background: lightcoral;
  border: none;
  font-family: "Courier New", Courier, monospace;
  font-weight: 600;
`;

const MiniHeader = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: space-evenly;
  align-content: center;
  align-items: center;
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

const ContentHolder = styled.div`
  padding: 20px;
  border-left: 3px solid black;
  margin: 15px;
`;

const MainContent = styled.div`
  max-width: 900px;
`;

const Subscriptions = () => {
  const [npToken, setnpToken] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [dbdata, setdbData] = useState([]);
  const [skip, setSkip] = useState(0);

  let take = 5;

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
    setLoading(false);
  };

  const fetchFromDB = async () => {
    setLoading(true);
    let newskip = skip + take;
    setSkip(newskip);
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

  const handleDBRefetch = async () => {
    fetchFromDB();
  };

  return (
    <Wrapper>
      <MiniHeader>
        <Title> Manage your Subscriptions </Title>
        <div
          style={{
            display: "grid",
            gridGap: "30px",
            gridTemplateColumns: "2fr 1fr"
          }}
        >
          <FetchMoreButton onClick={handleFetchMore}>
            Fetch More From API
          </FetchMoreButton>

          <div>
            {isLoading ? (
              <div style={{ display: "grid" }}>
                {" "}
                <LoadingLogo /> <span> Loading ... </span>{" "}
              </div>
            ) : (
              <span></span>
            )}
          </div>
        </div>
      </MiniHeader>

      <MainContent>
        {dbdata.map((x, i) => (
          <div key={i}>
            {x.map((y, i) => (
              <ContentHolder key={i}>
                {" "}
                {y.sender}
                {y.link !== "" ? (
                  <a href={y.link}> Unsub </a>
                ) : (
                  <a href={y.link}> Missing Link </a>
                )}
                <p>
                  {" "}
                  <strong> Latest Contact: </strong> {y.date}
                </p>
              </ContentHolder>
            ))}
          </div>
        ))}

        <SeeMoreButton onClick={handleDBRefetch}> See More </SeeMoreButton>
      </MainContent>
    </Wrapper>
  );
};

export default Subscriptions;
