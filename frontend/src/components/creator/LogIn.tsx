const SignUp: React.FC = () => {
    return (
        <form className="flex flex-col space-y-4">
            <input
                type="text"
                placeholder="Full Name"
                className="input input-bordered w-full"
            />
            <input
                type="email"
                placeholder="Email"
                className="input input-bordered w-full"
            />
            <input
                type="password"
                placeholder="Password"
                className="input input-bordered w-full"
            />
            <button type="submit" className="btn btn-primary w-full">
                Sign Up
            </button>
        </form>
    );
};

export default SignUp;
