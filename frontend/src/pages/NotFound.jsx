import React from "react";
import styled from "styled-components";
import Header from "../components/Header";
import Footer from "../components/Footer";

const NotFound = () => {
  const NotFound = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    font-size: 3rem;
  `;

  return (
    <>
      <Header />
      <NotFound>404 Page Not Found</NotFound>
      <Footer />
    </>
  );
};

export default NotFound;
