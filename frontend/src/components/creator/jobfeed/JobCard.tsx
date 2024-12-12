import React from "react";

type JobCardProps = {
    title: string;
    description: string;
    locations: string[];
    companies: string[];
    jobs: string[];
    date: string;
};

const JobCard: React.FC<JobCardProps> = ({ title, description, locations, companies, jobs, date }) => {
    return (
        <div className="bg-white p-6 rounded-lg shadow hover:shadow-lg transition">
            <h3 className="text-lg font-medium text-gray-800 mb-2">
                {title} <span className="text-purple-500 text-sm">(Why I see this?)</span>
            </h3>
            <p className="text-gray-600 mb-4">{description}</p>
            <div className="flex flex-wrap gap-2 mb-4">
                {locations.map((loc, idx) => (
                    <span
                        key={idx}
                        className="bg-green-100 text-green-600 text-sm px-3 py-1 rounded-full"
                    >
                        {loc}
                    </span>
                ))}
            </div>
            <div className="flex flex-wrap gap-2 mb-4">
                {companies.map((comp, idx) => (
                    <span
                        key={idx}
                        className="bg-blue-100 text-blue-600 text-sm px-3 py-1 rounded-full"
                    >
                        {comp}
                    </span>
                ))}
            </div>
            <h4 className="font-medium text-gray-800 mb-2">Collection job preview:</h4>
            <ul className="list-disc list-inside text-gray-600 mb-4">
                {jobs.map((job, idx) => (
                    <li key={idx}>{job}</li>
                ))}
            </ul>
            <div className="flex justify-between items-center">
                <span className="text-sm text-gray-500">{date}</span>
                <button className="bg-purple-500 text-white px-4 py-2 rounded-md hover:bg-purple-600 transition">
                    Subscribe
                </button>
            </div>
        </div>
    );
};

export default JobCard;
