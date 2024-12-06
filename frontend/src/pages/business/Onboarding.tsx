import React, { useState } from "react";

const BusinessOnboarding: React.FC = () => {
  const [formData, setFormData] = useState({
    email: "user@example.com", // Hardcoded email
    phoneNumber: "123-456-7890", // Hardcoded phone number
    firstName: "",
    lastName: "",
    businessName: "",
    businessType: "",
    introduction: "",
    address: "",
    zipcode: "",
    city: "",
    state: "",
    country: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { id, value } = e.target;
    setFormData((prev) => ({ ...prev, [id]: value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Submitted data:", formData);
    // Add form submission logic here
  };

  return (
    <div className="flex min-h-screen mb-5">
      {/* Left Panel */}
      <div className="w-1/2 flex items-center justify-center">
        <div className="text-left">
          <h1 className="text-2xl font-bold mb-4">Onboarding as a Business</h1>
          <p className="text-gray-700">
            Create your business's profile to find your best content creators!
          </p>
        </div>
      </div>

      {/* Right Panel */}
      <div className="w-1/2 flex items-center justify-center">
        <form onSubmit={handleSubmit} className="space-y-6 max-w-lg w-full">
          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="email">
              Email
            </label>
            <input
              type="email"
              id="email"
              value={formData.email}
              readOnly
              className="w-full border border-gray-300 rounded px-3 py-2 text-gray-700"
            />
          </div>

          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="phoneNumber">
              Phone Number
            </label>
            <input
              type="text"
              id="phoneNumber"
              value={formData.phoneNumber}
              readOnly
              className="w-full border border-gray-300 rounded px-3 py-2 text-gray-700"
            />
          </div>

          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="firstName">
              First Name
            </label>
            <input
              type="text"
              id="firstName"
              value={formData.firstName}
              onChange={handleChange}
              className="w-full border border-gray-300 rounded px-3 py-2"
            />
          </div>

          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="lastName">
              Last Name
            </label>
            <input
              type="text"
              id="lastName"
              value={formData.lastName}
              onChange={handleChange}
              className="w-full border border-gray-300 rounded px-3 py-2"
            />
          </div>

          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="businessName">
              Business Name
            </label>
            <input
              type="text"
              id="businessName"
              value={formData.businessName}
              onChange={handleChange}
              className="w-full border border-gray-300 rounded px-3 py-2"
            />
          </div>

          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="businessType">
              Business Type
            </label>
            <input
              type="text"
              id="businessType"
              value={formData.businessType}
              onChange={handleChange}
              className="w-full border border-gray-300 rounded px-3 py-2"
            />
          </div>


          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="introduction">
              Introduction
            </label>
            <textarea
              id="introduction"
              value={formData.introduction}
              onChange={handleChange}
              className="w-full border border-gray-300 rounded px-3 py-2"
              rows={4}
            ></textarea>
          </div>

          <div>
            <label className="block text-gray-700 font-bold mb-1" htmlFor="address">
              Address
            </label>
            <input
              type="text"
              id="address"
              value={formData.address}
              onChange={handleChange}
              className="w-full border border-gray-300 rounded px-3 py-2"
            />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-gray-700 font-bold mb-1" htmlFor="zipcode">
                Zipcode
              </label>
              <input
                type="text"
                id="zipcode"
                value={formData.zipcode}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded px-3 py-2"
              />
            </div>
            <div>
              <label className="block text-gray-700 font-bold mb-1" htmlFor="city">
                City
              </label>
              <input
                type="text"
                id="city"
                value={formData.city}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded px-3 py-2"
              />
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-gray-700 font-bold mb-1" htmlFor="state">
                State
              </label>
              <input
                type="text"
                id="state"
                value={formData.state}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded px-3 py-2"
              />
            </div>
            <div>
              <label className="block text-gray-700 font-bold mb-1" htmlFor="country">
                Country
              </label>
              <input
                type="text"
                id="country"
                value={formData.country}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded px-3 py-2"
              />
            </div>
          </div>

          <div>
            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
            >
              Submit
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default BusinessOnboarding;

