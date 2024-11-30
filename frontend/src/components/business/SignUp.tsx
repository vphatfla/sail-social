import React, { useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';

const SignUp: React.FC = () => {
    // default values for testing
    const defaultEmail = "example123@gmail.com"
    const defaultPhone = "123-456-7891"
    const defaultPassword = "123456"

    const url = import.meta.env.VITE_SERVER_URL
    const [email, setEmail] = useState(defaultEmail);
    const [phone, setPhone] = useState(defaultPhone);
    const [password, setPassword] = useState(defaultPassword);
    const [confirmPassword, setConfirmPassword] = useState(defaultPassword);
    const [error, setError] = useState('');
    const navigate = useNavigate();



    // Validation checks
    const isEmailValid = email.includes('@');
    const isPhoneValid = phone.length === 12; // Auto-formatted to xxx-xxx-xxxx
    const isPasswordValid = password.length >= 6;
    const isPasswordMatching = password === confirmPassword;

    const isFormValid = useMemo(() => {
        const isEmailValid = email.includes('@');
        const isPhoneValid = phone.length === 12; // xxx-xxx-xxxx
        const isPasswordValid = password.length >= 6;
        const isPasswordMatching = password === confirmPassword;
        return isEmailValid && isPhoneValid && isPasswordValid && isPasswordMatching;
    }, [email, phone, password, confirmPassword]);

    // Auto-format phone number
    const handlePhoneChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        let value = e.target.value.replace(/\D/g, ''); // Remove non-numeric characters
        if (value.length > 3 && value.length <= 6) {
            value = `${value.slice(0, 3)}-${value.slice(3)}`;
        } else if (value.length > 6) {
            value = `${value.slice(0, 3)}-${value.slice(3, 6)}-${value.slice(6, 10)}`;
        }
        setPhone(value);
    };

    // Handle form submission
    const handleSignUp = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(''); // Clear any previous errors

        try {
            const response = await fetch(url + '/api/creator/sign-up', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, phone, password }),
            });

            if (!response.ok) {
                const data = await response.json();
                throw new Error(data.error || 'Something went wrong.');
            }

            navigate('/creator'); // Redirect to Creator page on success
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
        } catch (err: any) {
            setError(err.message);
        }
    };

    return (
        <form onSubmit={handleSignUp} className="max-w-md mx-auto">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="email"
                >
                    Email
                </label>
                <input
                    className={`shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${isEmailValid || !email ? '' : 'border-red-500'
                        }`}
                    id="email"
                    type="email"
                    placeholder="example123@gmail.com"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                {email && !isEmailValid && (
                    <p className="text-red-500 text-sm">Email must include "@"</p>
                )}
            </div>

            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="phone"
                >
                    Phone Number
                </label>
                <input
                    className={`shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${isPhoneValid || !phone ? '' : 'border-red-500'
                        }`}
                    id="phone"
                    type="text"
                    placeholder="123-456-7891"
                    value={phone}
                    onChange={handlePhoneChange}
                />
                {phone && !isPhoneValid && (
                    <p className="text-red-500 text-sm">Phone number must be valid.</p>
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
                    className={`shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${isPasswordValid || !password ? '' : 'border-red-500'
                        }`}
                    id="password"
                    type="password"
                    placeholder="123456"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                {password && !isPasswordValid && (
                    <p className="text-red-500 text-sm">
                        Password must be at least 6 characters.
                    </p>
                )}
            </div>

            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="confirmPassword"
                >
                    Retype Password
                </label>
                <input
                    className={`shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline ${isPasswordMatching || !confirmPassword ? '' : 'border-red-500'
                        }`}
                    id="confirmPassword"
                    type="password"
                    placeholder="13456"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                />
                {confirmPassword && !isPasswordMatching && (
                    <p className="text-red-500 text-sm">Passwords do not match.</p>
                )}
            </div>

            {error && <p className="text-red-500 text-sm mb-4">{error}</p>}

            <div className="mb-4">
                <button
                    type="submit"
                    className={isFormValid ? "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full" : "bg-blue-500 text-white font-bold py-2 px-4 rounded opacity-50 cursor-not-allowed w-full"}
                    disabled={!isFormValid}
                >
                    Sign Up
                </button>
            </div>
        </form>
    );
};

export default SignUp;
