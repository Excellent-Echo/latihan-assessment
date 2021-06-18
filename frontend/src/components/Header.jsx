import React, { useEffect, useState } from "react";
import { Navbar, Nav } from "react-bootstrap";
import { useLocation } from "react-router-dom";

const Header = () => {
  const [pageURL, setPageURL] = useState("");
  const location = useLocation();
  useEffect(() => {
    setPageURL(location.pathname);
  }, []);

  return (
    <>
      <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
        <Navbar.Brand href="/" className="ml-5">
          BOOK LIST APP
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse
          id="responsive-navbar-nav"
          className="justify-content-end"
        >
          <Nav>
            {pageURL === "/register" || (
              <Nav.Link href="/register" className="mr-5">
                Register
              </Nav.Link>
            )}
            {pageURL === "/login" || (
              <Nav.Link href="/login" className="mr-5">
                Login
              </Nav.Link>
            )}
          </Nav>
        </Navbar.Collapse>
      </Navbar>
    </>
  );
};

export default Header;
