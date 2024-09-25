import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

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
    <nav
      className={`w-full fixed top-0 z-50 transition-all duration-300 ${
        scrolling ? 'bg-white shadow-lg' : 'bg-transparent'
      }`}
    >
      <div className="container mx-auto py-3 flex justify-between items-center">
        <div className='flex items-center'>
          <Link to="/" className="text-2xl font-bold text-black">
            Booster Hub
          </Link>
          <div className="hidden md:flex space-x-4 ml-8">
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
        <div className="flex items-center space-x-4">
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
    </nav>
  );
};

export default NavBar;
