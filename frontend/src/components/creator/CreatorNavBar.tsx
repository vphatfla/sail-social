import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { FaChevronDown } from 'react-icons/fa';

const CreatorNavBar: React.FC = () => {
    const [dropdownOpen, setDropdownOpen] = useState(false);

    return (
        <nav className="w-full fixed top-0 z-50 bg-white">
            <div className="container mx-auto py-3 flex justify-between items-center">
                {/* Left: Logo */}
                <span className="text-2xl font-bold text-black">
                    Sail Social - Creator
                </span>

                {/* Middle: Navigation Links */}
                <div className="flex space-x-4">
                    <NavLink
                        to="/creator/job-feed"
                        className={({ isActive }) =>
                            `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'}`
                        }
                    >
                        Job Feed
                    </NavLink>
                    <NavLink
                        to="/creator/applied-jobs"
                        className={({ isActive }) =>
                            `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'}`
                        }
                    >
                        Applied Jobs
                    </NavLink>
                    <NavLink
                        to="/creator/notifications"
                        className={({ isActive }) =>
                            `transition-colors ${isActive ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-700 hover:text-blue-500'}`
                        }
                    >
                        Notifications
                    </NavLink>
                </div>

                {/* Right: Account Dropdown */}
                <div className="relative">
                    <button
                        onClick={() => setDropdownOpen(!dropdownOpen)}
                        className="flex items-center space-x-2 text-gray-700 hover:text-blue-500 transition-colors"
                    >
                        Account <FaChevronDown />
                    </button>

                    {dropdownOpen && (
                        <div className="absolute right-0 mt-2 w-48 bg-white border rounded-lg shadow-lg">
                            <button className="block w-full text-left px-4 py-2 hover:bg-gray-100">
                                Profile
                            </button>
                            <button className="block w-full text-left px-4 py-2 hover:bg-gray-100">
                                Settings
                            </button>
                            <button className="block w-full text-left px-4 py-2 hover:bg-gray-100">
                                Logout
                            </button>
                        </div>
                    )}
                </div>
            </div>
        </nav>
    );
};

export default CreatorNavBar;