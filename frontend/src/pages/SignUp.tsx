import React, { useState } from 'react';
import { Form, Button, ToggleButtonGroup, ToggleButton } from 'react-bootstrap';
import {motion} from "framer-motion"
import { useAuth } from '../contexts/AuthContext';
import { decryptToken } from '../utils/TokenUtils';
import { useNavigate } from 'react-router-dom';
import { callApi } from '../utils/ApiUtils';

const SignUp: React.FC = () => {
  const [userType, setUserType] = useState<string>('creator');
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);
  const { setUserInfo, isAuthenticated } = useAuth();
  const navigate = useNavigate();

  if (isAuthenticated()){
    navigate("/")
  }

  const handleToggleChange = (val: string) => {
    setUserType(val);
    setErrorMessage(null);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    setErrorMessage(null);
    setSuccessMessage(null);
  
    const formData = new FormData(event.target as HTMLFormElement);
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;
    const confirmPassword = formData.get('confirmPassword') as string;
    const phoneNumber = formData.get('phoneNumber') as string;
    const businessName = formData.get('businessName') as string | null;
  
    if (password !== confirmPassword) {
      setErrorMessage('Passwords do not match!');
      return;
    }
  
    const data = {
      email,
      password,
      phoneNumber,
      ...(userType === 'business' && { businessName }),
    };
    
    const endpoint = userType === 'creator'
    ? '/creator/sign-up'
    : '/business/sign-up';

    const response = await callApi(endpoint, "POST", data)

    if (response.error){
      setErrorMessage(response.error || 'Login failed, please try again.');
    }
    else {
      const token = response.data.token;
      localStorage.setItem('BoosterHubToken', token);

      const decodedInfo = decryptToken(token);
      setUserInfo(decodedInfo);

      navigate('/');
    }
  };
  
  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
    >
      <h2 className="text-center">Sign Up</h2>

      <ToggleButtonGroup
        type="radio"
        name="userType"
        value={userType}
        onChange={handleToggleChange}
        className="mb-4 d-flex justify-content-center"
      >
        <ToggleButton id="creator-toggle" value="creator" variant="outline-primary">
          Creator
        </ToggleButton>
        <ToggleButton id="business-toggle" value="business" variant="outline-primary">
          Local Business
        </ToggleButton>
      </ToggleButtonGroup>

      <Form onSubmit={handleSubmit} className="signup-form">
        <Form.Group controlId="formEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control type="email" name="email" placeholder="Enter email" required />
        </Form.Group>

        <Form.Group controlId="formPassword" className="mt-3">
          <Form.Label>Password</Form.Label>
          <Form.Control type="password" name="password" placeholder="Password" required />
        </Form.Group>

        <Form.Group controlId="formConfirmPassword" className="mt-3">
          <Form.Label>Confirm Password</Form.Label>
          <Form.Control type="password" name="confirmPassword" placeholder="Confirm Password" required />
        </Form.Group>

        <Form.Group controlId="formPhoneNumber" className="mt-3">
          <Form.Label>Phone Number (optional)</Form.Label>
          <Form.Control type="tel" name="phoneNumber" placeholder="Phone Number" />
        </Form.Group>

        {errorMessage && <p className="text-danger">{errorMessage}</p>}
        {successMessage && <p className="text-success">{successMessage}</p>}

        <Button variant="primary" type="submit" className="mt-4">
          Sign Up as {userType === 'creator' ? 'Creator' : 'Local Business'}
        </Button>
      </Form>
    </motion.div>
  );
};

export default SignUp;
