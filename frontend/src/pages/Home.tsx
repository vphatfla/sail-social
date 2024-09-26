import React from 'react';
import {motion} from "framer-motion"

const Home: React.FC = () => {
  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
    >
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
      <p>
        Build, manage and measure your creator community with Booster Hub's platform and expert support team.
      </p>
      <button className="btn btn-primary">I'M A BRAND</button>
      <button className="btn btn-outline-primary">I'M A CREATOR</button>
    </motion.div>
  );
};

export default Home;
