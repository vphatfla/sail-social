import React, { useState } from "react";

const LoginForm: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isEmailValid, setIsEmailValid] = useState(true);
  const [error, setError] = useState("");

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    if (!email.includes("@")) {
      setIsEmailValid(false);
      setError("Invalid email address.");
      return;
    }
    setIsEmailValid(true);
    setError("");

    // Handle login logic here
    console.log("Logged in with:", { email, password });
  };

  const isFormValid = email && password && isEmailValid;

  return (
    <form onSubmit={handleLogin} className="max-w-md mx-auto">
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="email"
        >
          Email
        </label>
        <input
          className={`shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${
            isEmailValid || !email ? "" : "border-red-500"
          }`}
          id="email"
          type="email"
          placeholder="example123@gmail.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        {!isEmailValid && (
          <p className="text-red-500 text-sm">Email must include "@".</p>
        )}
      </div>

      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="password"
        >
          Password
        </label>
        <input
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="password"
          type="password"
          placeholder="Your password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>

      {error && <p className="text-red-500 text-sm mb-4">{error}</p>}

      <div className="mb-4">
        <button
          type="submit"
          className={`${
            isFormValid
              ? "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
              : "bg-blue-500 text-white font-bold py-2 px-4 rounded opacity-50 cursor-not-allowed w-full"
          }`}
          disabled={!isFormValid}
        >
          Log In
        </button>
      </div>
    </form>
  );
};

export default LoginForm;

