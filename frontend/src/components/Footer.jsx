import React from "react";
import styled from "styled-components";

const Footer = () => {
  const Footer = styled.footer`
    position: absolute;
    bottom: 0;
    text-align: center;
    background-color: #343a40;
    color: white;
    width: 100%;
  `;

  return (
    <>
      <Footer className="page-footer">
        <div className="footer-copyright text-center py-3">
          Impact Byte ©2021
        </div>
      </Footer>
    </>
  );
};

export default Footer;
