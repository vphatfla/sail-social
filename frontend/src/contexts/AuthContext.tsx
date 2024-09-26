import React, { createContext, useState, useEffect, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { decryptToken, isTokenExpired } from "../utils/TokenUtils";

interface AuthContextType {
  userInfo: any;
  setUserInfo: (userInfo: any) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [userInfo, setUserInfo] = useState<any>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      if (isTokenExpired(token)) {
        logout();
      } else {
        const decodedInfo = decryptToken(token);
        if (decodedInfo) {
          setUserInfo(decodedInfo);
        }
      }
    }
  }, []);

  const logout = () => {
    localStorage.removeItem('token');
    setUserInfo(null);
    navigate('/login');
  };

  return (
    <AuthContext.Provider value={{ userInfo, setUserInfo, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
