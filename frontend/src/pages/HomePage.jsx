import React from "react";
import { Button } from "react-bootstrap";
import { useHistory } from "react-router";

const HomePage = () =>{
    const history = useHistory();
    return (
        <>
            <Navbar>
  <Container>
    <Navbar.Brand href="#home">Navbar with text</Navbar.Brand>
    <Navbar.Toggle />
    <Navbar.Collapse className="justify-content-end">
      <Button onClick={e =>{
          e.preventDefault();
          history.push("/login")
      }}>
      </Button>
      <Button onClick={e =>{
          e.preventDefault();
          history.push("/register")
      }}>
      </Button>
    </Navbar.Collapse>
  </Container>
</Navbar>
        </>
    )
}

export default HomePage;