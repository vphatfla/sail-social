import React, { useState, useEffect } from 'react';
import { Navbar, Nav, Container, Button } from 'react-bootstrap';
import { LinkContainer } from 'react-router-bootstrap';
import './NavBar.css';

const NavBar: React.FC = () => {
  const [scrolling, setScrolling] = useState(false);

  const handleScroll = () => {
    if (window.scrollY > 50) {
      setScrolling(true);
    } else {
      setScrolling(false);
    }
  };

  useEffect(() => {
    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, []);

  return (
    <Navbar
      expand="lg"
      className={`navbar-custom ${scrolling ? 'navbar-scrolled' : ''}`}
      fixed="top"
    >
      <Container>
        <LinkContainer to="/">
          <Navbar.Brand>Booster Hub</Navbar.Brand>
        </LinkContainer>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <LinkContainer to="/learn">
              <Nav.Link>Learn</Nav.Link>
            </LinkContainer>
            <LinkContainer to="/business">
              <Nav.Link>Business</Nav.Link>
            </LinkContainer>
            <LinkContainer to="/creators">
              <Nav.Link>Creators</Nav.Link>
            </LinkContainer>
            <LinkContainer to="/pricing">
              <Nav.Link>Pricing</Nav.Link>
            </LinkContainer>
          </Nav>
          <LinkContainer to="/login">
            <Nav.Link className="login-button">Login</Nav.Link>
          </LinkContainer>
          <LinkContainer to="/signup">
            <Button className="signup-button">Get in Touch</Button>
          </LinkContainer>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};

export default NavBar;
