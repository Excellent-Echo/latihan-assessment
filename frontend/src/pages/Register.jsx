import React, { useEffect } from "react";
import { Form, Button, Alert } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import userRegisterAction from "../redux/user/register/userRegisterAction";
import Header from "../components/Header";
import Footer from "../components/Footer";
import styled from "styled-components";

const Register = () => {
  const userRegisterData = useSelector((state) => state.userRegister);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(userRegisterAction.resetForm());
  }, []);

  const handleRegisterSubmit = (e) => {
    e.preventDefault();
    dispatch(
      userRegisterAction.register(
        userRegisterData.name,
        userRegisterData.address,
        userRegisterData.dateBirth,
        userRegisterData.email,
        userRegisterData.password
      )
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
            {userRegisterData.errorMessage && (
              <Alert variant="danger" dismissible>
                <Alert.Heading>{userRegisterData.errorMessage}</Alert.Heading>
              </Alert>
            )}
            {userRegisterData.successMessage && (
              <Alert variant="success" dismissible>
                <Alert.Heading>{userRegisterData.successMessage}</Alert.Heading>
              </Alert>
            )}
            {!userRegisterData.successMessage && (
              <Form
                onSubmit={handleRegisterSubmit}
                className="d-flex flex-column justify-content-center align-items-center"
              >
                <Form.Group controlId="formBasicText" className="p-2 w-100">
                  <Form.Control
                    type="text"
                    placeholder="Name"
                    required
                    value={userRegisterData.name}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setName(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicText" className="p-2 w-100">
                  <Form.Control
                    type="text"
                    placeholder="Address"
                    required
                    value={userRegisterData.address}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setAddress(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="date" className="p-2 w-100">
                  <Form.Control
                    type="date"
                    required
                    placeholder="Birth Date"
                    value={userRegisterData.dateBirth}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setDateBirth(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicEmail" className="p-2 w-100">
                  <Form.Control
                    type="email"
                    placeholder="Email"
                    required
                    value={userRegisterData.email}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setEmail(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicPassword" className="p-2 w-100">
                  <Form.Control
                    type="password"
                    placeholder="Password"
                    required
                    autocomplete="new-password"
                    value={userRegisterData.password}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setPassword(e.target.value))
                    }
                  />
                </Form.Group>
                <Button
                  variant="primary"
                  type="submit"
                  disabled={userRegisterData.isLoading ? true : false}
                >
                  {userRegisterData.isLoading ? "Loading..." : "Register"}
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

export default Register;
