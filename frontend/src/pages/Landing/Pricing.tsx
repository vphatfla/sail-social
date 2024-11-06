import React from 'react';
import { motion } from "framer-motion"

const Pricing: React.FC = () => {
  return (
    <motion.div
      initial={{ opacity: 0, x: -100 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 100 }}
      transition={{ duration: 0.5 }}
    >
      <h1>Pricing</h1>
      <p>
        Explore our pricing plans that suit businesses and creators at every stage of their journey.
      </p>
    </motion.div>
  );
};

export default Pricing;
