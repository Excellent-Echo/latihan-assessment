import React, { useEffect } from "react";
import { Form, Button, Alert } from "react-bootstrap";
import { useDispatch, useSelector } from "react-redux";
import userRegisterAction from "../redux/user/register/userRegisterAction";

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
      <div className="container">
        <div className="row">
          <div className="col">
            <h1 className="text-center">Register</h1>
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
              <Form onSubmit={handleRegisterSubmit}>
                <Form.Group controlId="formBasicText">
                  <Form.Label>Name</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="Enter name"
                    required
                    value={userRegisterData.name}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setName(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicText">
                  <Form.Label>Address</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="Enter address"
                    required
                    value={userRegisterData.address}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setAddress(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="date">
                  <Form.Label>Birth Date</Form.Label>
                  <Form.Control
                    type="date"
                    required
                    value={userRegisterData.dateBirth}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setDateBirth(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicEmail">
                  <Form.Label>Email address</Form.Label>
                  <Form.Control
                    type="email"
                    placeholder="Enter email"
                    required
                    value={userRegisterData.email}
                    onChange={(e) =>
                      dispatch(userRegisterAction.setEmail(e.target.value))
                    }
                  />
                  <Form.Text className="text-muted"></Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicPassword">
                  <Form.Label>Password</Form.Label>
                  <Form.Control
                    type="password"
                    placeholder="Password"
                    required
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
                  {userRegisterData.isLoading ? "Loading..." : "Submit"}
                </Button>
              </Form>
            )}
          </div>
        </div>
      </div>
    </>
  );
};

export default Register;
