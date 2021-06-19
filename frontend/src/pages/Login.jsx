import React, { useEffect } from "react";
import { Form, Button, Alert } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import userRegisterAction from "../redux/user/register/userRegisterAction";
import Header from "../components/Header";
import Footer from "../components/Footer";
import styled from "styled-components";
import userLoginAction from "../redux/user/login/userLoginAction";

const Login = () => {
  const userLoginData = useSelector((state) => state.userLogin);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(userLoginAction.resetForm());
    console.log(userLoginData);
  }, []);

  const handleLoginSubmit = (e) => {
    e.preventDefault();
    dispatch(
      userLoginAction.login(userLoginData.email, userLoginData.password)
    );
  };

  return (
    <>
      <Header />
      <div className="container">
        <div className="row justify-content-center mt-5">
          <div
            className="col-sm-6"
            style={{ backgroundColor: "#eaeaea", padding: "2rem" }}
          >
            {userLoginData.errorMessage && (
              <Alert variant="danger" dismissible>
                <Alert.Heading>{userLoginData.errorMessage}</Alert.Heading>
              </Alert>
            )}
            {userLoginData.successMessage && (
              <Alert variant="success">
                <Alert.Heading>{userLoginData.successMessage}</Alert.Heading>
              </Alert>
            )}
            {!userLoginData.successMessage && (
              <Form
                onSubmit={handleLoginSubmit}
                className="d-flex flex-column justify-content-center align-items-center"
              >
                <Form.Group controlId="formBasicEmail" className="p-2 w-100">
                  <Form.Control
                    type="email"
                    placeholder="Email"
                    required
                    value={userLoginData.email}
                    onChange={(e) =>
                      dispatch(userLoginAction.setEmail(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicPassword" className="p-2 w-100">
                  <Form.Control
                    type="password"
                    placeholder="Password"
                    required
                    value={userLoginData.password}
                    onChange={(e) =>
                      dispatch(userLoginAction.setPassword(e.target.value))
                    }
                  />
                </Form.Group>
                <Button
                  variant="primary"
                  type="submit"
                  disabled={userLoginData.isLoading ? true : false}
                >
                  {userLoginData.isLoading ? "Loading..." : "Login"}
                </Button>
              </Form>
            )}
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Login;
