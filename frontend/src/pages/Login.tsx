import React, { useState } from 'react';
import { Form, Button, Container, ToggleButtonGroup, ToggleButton } from 'react-bootstrap';
import { useAuth } from '../contexts/AuthContext';
import { decryptToken } from '../utils/TokenUtils';
import { useNavigate } from 'react-router-dom';
import { callApi } from '../utils/ApiUtils';

const Login: React.FC = () => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [userType, setUserType] = useState<string>('creator');
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const { setUserInfo } = useAuth();
  const navigate = useNavigate();

  const handleToggleChange = (val: string) => {
    setUserType(val);
    setErrorMessage(null);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    setErrorMessage(null);

    const data = {
      email,
      password,
    };

    const endpoint = userType === 'creator'
      ? '/creator/log-in'
      : '/business/log-in';

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
      console.log('Login successful!');
    }
  };

  return (
    <Container className="page-container">
      <h2 className="text-center">Login</h2>

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

      <Form onSubmit={handleSubmit}>
        <Form.Group controlId="formEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            type="email"
            placeholder="Enter email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </Form.Group>

        <Form.Group controlId="formPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </Form.Group>

        {errorMessage && <p className="text-danger">{errorMessage}</p>}

        <Button variant="primary" type="submit" className="mt-4">
          Login as {userType === 'creator' ? 'Creator' : 'Local Business'}
        </Button>
      </Form>
    </Container>
  );
};

export default Login;
