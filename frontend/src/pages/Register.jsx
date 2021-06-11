import React, {useEffect} from "react"
import { useSelector, useDispatch } from "react-redux"
import userRegisterAction from "../redux/user/register/userRegisterAction"
import { Form } from "react-bootstrap"

const RegisterPage = () =>{
    const userRegisterData = useSelector(state => state.userRegister);
    const dispatch = useDispatch()

    useEffect(() =>{
        dispatch(userRegisterAction.resetForm())
    }, [])

    const handleRegisterSubmit = e =>{
        e.preventDefault();
        dispatch(
            userRegisterAction.register(
                userRegisterData.Name,
                userRegisterData.Email,
                userRegisterData.Password,
                userRegisterData.Address,
                userRegisterData.DateBirth
            )
        )
    }
    return (
        <div className="container">
            <Form>
  <Form.Group className="mb-3" controlId="formBasicName">
    <Form.Label>Name</Form.Label>
    <Form.Control type="name" placeholder="Enter Name" />
  </Form.Group>
  <Form.Group className="mb-3" controlId="formBasicAddress">
    <Form.Label>Address</Form.Label>
    <Form.Control type="name" placeholder="Enter Address" />
  </Form.Group>
  <Form.Group className="mb-3" controlId="formBasicBirth">
    <Form.Label>Date Birth</Form.Label>
    <Form.Control type="name" placeholder="Enter Date Birth" />
  </Form.Group>
  <Form.Group className="mb-3" controlId="formBasicEmail">
    <Form.Label>Email address</Form.Label>
    <Form.Control type="email" placeholder="Enter email" />
    <Form.Text className="text-muted">
      We'll never share your email with anyone else.
    </Form.Text>
  </Form.Group>

  <Form.Group className="mb-3" controlId="formBasicPassword">
    <Form.Label>Password</Form.Label>
    <Form.Control type="password" placeholder="Password" />
  </Form.Group>
  <Button variant="primary" type="submit">
    Submit
  </Button>
</Form>
</div>
    )
}




export default RegisterPage;
