import React from 'react';
import { Form, Button, Container } from 'react-bootstrap';

const SignUp: React.FC = () => {
  return (
    <Container className="signup-container">
      <h2 className="text-center">Sign Up</h2>
      <Form className="signup-form">
        <Form.Group controlId="formEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control type="email" placeholder="Enter email" required />
        </Form.Group>

        <Form.Group controlId="formPassword" className="mt-3">
          <Form.Label>Password</Form.Label>
          <Form.Control type="password" placeholder="Password" required />
        </Form.Group>

        <Form.Group controlId="formConfirmPassword" className="mt-3">
          <Form.Label>Confirm Password</Form.Label>
          <Form.Control type="password" placeholder="Confirm Password" required />
        </Form.Group>

        <Button variant="primary" type="submit" className="mt-4">
          Sign Up
        </Button>
      </Form>
    </Container>
  );
};

export default SignUp;
