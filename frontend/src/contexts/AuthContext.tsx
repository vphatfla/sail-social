// import React, { createContext, useState, useEffect, useContext } from 'react';
// import { useNavigate } from 'react-router-dom';
// import { decryptToken, isTokenExpired } from "../utils/TokenUtils";

// interface AuthContextType {
//   // eslint-disable-next-line @typescript-eslint/no-explicit-any
//   userInfo: any;
//   // eslint-disable-next-line @typescript-eslint/no-explicit-any
//   setUserInfo: (userInfo: any) => void;
//   logout: () => void;
//   isAuthenticated: () => boolean;
//   isOnboarded: () => boolean;
//   isCreator: () => boolean;
//   isBusiness: () => boolean;
// }

// const CREATOR: string = "creator";
// const BUSINESS: string = "business";

// const AuthContext = createContext<AuthContextType | undefined>(undefined);

// export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
//   // eslint-disable-next-line @typescript-eslint/no-explicit-any
//   const [userInfo, setUserInfo] = useState<any>(null);
//   const navigate = useNavigate();

//   useEffect(() => {
//     const token = localStorage.getItem('BoosterHubToken');
//     if (token) {
//       if (isTokenExpired(token)) {
//         logout();
//       } else {
//         const decodedInfo = decryptToken(token);
//         if (decodedInfo) {
//           setUserInfo(decodedInfo);
//         }
//       }
//     }
//   }, [setUserInfo]);

//   const logout = () => {
//     localStorage.removeItem('BoosterHubToken');
//     setUserInfo(null);
//     navigate('/login');
//   };

//   const isAuthenticated = () => {
//     return !!localStorage.getItem('BoosterHubToken') && userInfo != null;
//   };

//   const isOnboarded = () => {
//     return false
//   }

//   const isCreator = () => {
//     return userInfo["typeUser"] === CREATOR
//   }

//   const isBusiness = () => {
//     return userInfo["typeUser"] === BUSINESS
//   }

//   return (
//     <AuthContext.Provider value={{ userInfo, setUserInfo, logout, isAuthenticated, isOnboarded, isCreator, isBusiness }}>
//       {children}
//     </AuthContext.Provider>
//   );
// };

// export const useAuth = () => {
//   const context = useContext(AuthContext);
//   if (!context) {
//     throw new Error('useAuth must be used within an AuthProvider');
//   }
//   return context;
// };
