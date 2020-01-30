import React from "react";
import styled from "styled-components";
import { ReactComponent as FetchLogo } from "./fetch.svg";
import { ReactComponent as ParseLogo } from "./parse.svg";
import { ReactComponent as UnsubLogo } from "./unsub.svg";

const Wrapper = styled.div`
  padding: 30px;
`;
const Title = styled.div`
  font-size: 1.2em;
  text-align: center;
  color: palevioletred;
`;

const ContentWrapper = styled.div`
  display: flex;
  justify-content: space-around;
  margin-top: 60px;
  flex-wrap: wrap;
  font-family: "Courier New", Courier, monospace;
  font-weight: 700;
`;

const DescriptionWrapper = styled.div`
  display: grid;
  justify-items: center;
`;

const Para = styled.p`
  max-width: 250px;
`;

const Home = () => {
  return (
    <Wrapper>
      <Title>SmailTrail Works in 3 Simple steps</Title>
      <ContentWrapper>
        <DescriptionWrapper>
          <FetchLogo />
          <Para> Fetch your emails per page </Para>
        </DescriptionWrapper>

        <DescriptionWrapper>
          <ParseLogo />
          <Para>
            Only take what is required and since this app is run using your
            personal credentials.json No one else will ever see your data
          </Para>
        </DescriptionWrapper>
        <DescriptionWrapper>
          <UnsubLogo />
          <Para>
            {" "}
            After parsing the necessary elements you are provided with a one
            click unsubscribe option
          </Para>
        </DescriptionWrapper>
      </ContentWrapper>
    </Wrapper>
  );
};

export default Home;
