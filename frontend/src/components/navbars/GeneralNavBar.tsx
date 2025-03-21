import React, { useState, useEffect } from 'react';
import { NavLink } from 'react-router-dom';
import { FaBars, FaTimes } from 'react-icons/fa';

const GeneralNavBar: React.FC = () => {
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
      className={`w-full fixed top-0 z-50 transition-all duration-300 ${scrolling ? 'bg-white shadow-lg' : 'bg-transparent'
        }`}
    >
      <div className="container mx-auto py-3 flex justify-between items-center">
        <div className="flex items-center space-x-8">
          <NavLink
            to="/"
            className={({ isActive }) =>
              `text-2xl font-bold transition-colors ${(isActive) ? 'text-blue-500' : 'text-black hover:text-blue-500'
              }`
            }
          >
            Sail Social
          </NavLink>
          {
            (
              <div className="hidden md:flex space-x-4">
                <NavLink
                  to="/learn"
                  className={({ isActive }) =>
                    `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                    }`
                  }
                >
                  Learn
                </NavLink>
                <NavLink
                  to="/business"
                  className={({ isActive }) =>
                    `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                    }`
                  }
                >
                  Business
                </NavLink>
                <NavLink
                  to="/creators"
                  className={({ isActive }) =>
                    `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                    }`
                  }
                >
                  Creators
                </NavLink>
                <NavLink
                  to="/pricing"
                  className={({ isActive }) =>
                    `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                    }`
                  }
                >
                  Pricing
                </NavLink>
              </div>
            )
          }
        </div>

        {/* For phones */}
        <div className="md:hidden">
          <FaBars className="text-2xl cursor-pointer" onClick={toggleMenu} />
        </div>

        {/* For bigger screens */}
        <div className="hidden md:flex items-center space-x-4">

        </div>
      </div>

      {/* For smaller screens */}
      <div
        className={`fixed top-0 right-0 w-full h-full bg-white flex flex-col justify-center items-center space-y-6 transform ${menuOpen ? 'translate-x-0' : 'translate-x-full'
          } transition-transform duration-500 ease-in-out`}
      >
        <FaTimes className="absolute top-5 right-5 text-3xl cursor-pointer" onClick={toggleMenu} />
        {
          (
            <NavLink
              to="/learn"
              className={({ isActive }) =>
                `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                }`
              }
              onClick={toggleMenu}
            >
              Learn
            </NavLink>
          )
        }
        {
          (
            <NavLink
              to="/business"
              className={({ isActive }) =>
                `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                }`
              }
              onClick={toggleMenu}
            >
              Business
            </NavLink>
          )
        }
        {
          (
            <NavLink
              to="/creators"
              className={({ isActive }) =>
                `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                }`
              }
              onClick={toggleMenu}
            >
              Creators
            </NavLink>
          )
        }
        {
          (
            <NavLink
              to="/pricing"
              className={({ isActive }) =>
                `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'
                }`
              }
              onClick={toggleMenu}
            >
              Pricing
            </NavLink>
          )
        }


      </div>
    </nav>
  );
};

export default GeneralNavBar;
