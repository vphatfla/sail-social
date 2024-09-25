import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FaBars, FaTimes } from 'react-icons/fa';

const NavBar: React.FC = () => {
  const [scrolling, setScrolling] = useState(false);
  const [menuOpen, setMenuOpen] = useState(false);

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

  const toggleMenu = () => {
    setMenuOpen(!menuOpen);
  };

  return (
    <nav
      className={`w-full fixed top-0 z-50 transition-all duration-300 ${
        scrolling ? 'bg-white shadow-lg' : 'bg-transparent'
      }`}
    >
      <div className="container mx-auto py-3 flex justify-between items-center">
        {/* Grouping BoosterHub and Links together */}
        <div className="flex items-center space-x-8">
          <Link to="/" className="text-2xl font-bold text-black">
            Booster Hub
          </Link>
          {/* Desktop Links (hidden on mobile) */}
          <div className="hidden md:flex space-x-4">
            <Link to="/learn" className="text-gray-700 hover:text-black transition-colors">
              Learn
            </Link>
            <Link to="/business" className="text-gray-700 hover:text-black transition-colors">
              Business
            </Link>
            <Link to="/creators" className="text-gray-700 hover:text-black transition-colors">
              Creators
            </Link>
            <Link to="/pricing" className="text-gray-700 hover:text-black transition-colors">
              Pricing
            </Link>
          </div>
        </div>
        {/* Mobile Hamburger Menu Icon */}
        <div className="md:hidden">
          <FaBars className="text-2xl cursor-pointer" onClick={toggleMenu} />
        </div>
        {/* Login and Sign Up buttons (hidden on mobile) */}
        <div className="hidden md:flex items-center space-x-4">
          <Link to="/login" className="text-gray-700 hover:text-black transition-colors">
            Login
          </Link>
          <Link to="/signup">
            <button className="bg-blue-500 text-white px-4 py-2 rounded-full hover:bg-blue-600 transition-colors">
              Get in Touch
            </button>
          </Link>
        </div>
      </div>

      {/* Full-Screen Mobile Menu (Slide from right) */}
      <div
        className={`fixed top-0 right-0 w-full h-full bg-white flex flex-col justify-center items-center space-y-6 transform ${
          menuOpen ? 'translate-x-0' : 'translate-x-full'
        } transition-transform duration-500 ease-in-out`}
      >
        <FaTimes className="absolute top-5 right-5 text-3xl cursor-pointer" onClick={toggleMenu} />
        <Link to="/learn" className="text-xl text-gray-700 hover:text-black" onClick={toggleMenu}>
          Learn
        </Link>
        <Link to="/business" className="text-xl text-gray-700 hover:text-black" onClick={toggleMenu}>
          Business
        </Link>
        <Link to="/creators" className="text-xl text-gray-700 hover:text-black" onClick={toggleMenu}>
          Creators
        </Link>
        <Link to="/pricing" className="text-xl text-gray-700 hover:text-black" onClick={toggleMenu}>
          Pricing
        </Link>
        <Link to="/login" className="text-xl text-gray-700 hover:text-black" onClick={toggleMenu}>
          Login
        </Link>
        <Link to="/signup" onClick={toggleMenu}>
          <button className="bg-blue-500 text-white px-4 py-2 rounded-full hover:bg-blue-600 transition-colors">
            Get in Touch
          </button>
        </Link>
      </div>
    </nav>
  );
};

export default NavBar;
