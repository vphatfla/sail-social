import React from "react";
import { Outlet } from "react-router-dom";

const CreatorHome: React.FC = () => {
    return (
        <div className="mt-16 p-4">
            <Outlet />
        </div>
    );
};

export default CreatorHome;