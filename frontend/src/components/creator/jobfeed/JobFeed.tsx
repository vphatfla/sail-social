import React, { useState } from "react";
import JobCard from "./JobCard"; // Import the JobCard component

type JobCollection = {
    title: string;
    description: string;
    locations: string[];
    companies: string[];
    jobs: string[];
    date: string;
};

const collections: JobCollection[] = [
    {
        title: "Trainee Engineer",
        description: "This is a collection mini description. Maybe.",
        locations: ["Amsterdam", "San Francisco"],
        companies: ["McDonald's", "eBay", "Facebook", "IBM", "The Walt Disney Company"],
        jobs: ["Senior Software Developer", "Software Tester"],
        date: "Jan 18th",
    },
    {
        title: "Programmer Analyst",
        description: "This is a collection mini description. Maybe.",
        locations: ["Amsterdam", "San Francisco"],
        companies: ["McDonald's", "eBay", "Facebook", "IBM", "The Walt Disney Company"],
        jobs: ["Senior Software Developer", "Software Tester"],
        date: "Jan 18th",
    },
    {
        title: "QA Engineer",
        description: "This is a collection mini description. Maybe.",
        locations: ["Amsterdam", "San Francisco"],
        companies: ["McDonald's", "eBay", "Facebook", "IBM", "The Walt Disney Company"],
        jobs: ["Senior Software Developer", "Software Tester"],
        date: "Jan 18th",
    },
];

const JobFeedComponent: React.FC = () => {
    const [selectedRoles, setSelectedRoles] = useState<string[]>(["Software Engineer", "Software Tester"]);
    const [selectedIndustries, setSelectedIndustries] = useState<string[]>(["SaaS", "Machine learning"]);

    const handleRoleToggle = (role: string) => {
        setSelectedRoles((prev) =>
            prev.includes(role) ? prev.filter((r) => r !== role) : [...prev, role]
        );
    };

    const handleIndustryToggle = (industry: string) => {
        setSelectedIndustries((prev) =>
            prev.includes(industry) ? prev.filter((i) => i !== industry) : [...prev, industry]
        );
    };

    return (
        <div className="flex gap-4 p-6 bg-gray-100 min-h-screen">
            {/* Filters Section */}
            <aside className="w-64 bg-white rounded-lg shadow p-4">
                <h3 className="text-xl font-semibold mb-4">Job filters</h3>

                <div className="mb-6">
                    <h4 className="font-medium mb-2">Job availability</h4>
                    <div className="space-y-2">
                        <label className="flex items-center">
                            <input type="checkbox" className="mr-2" /> Open
                        </label>
                        <label className="flex items-center">
                            <input type="checkbox" className="mr-2" /> Closed
                        </label>
                    </div>
                </div>

                <div className="mb-6">
                    <h4 className="font-medium mb-2">Role</h4>
                    {["Software Engineer", "Software Tester", "QA Engineer", "Trainee"].map((role) => (
                        <label key={role} className="flex items-center mb-2">
                            <input
                                type="checkbox"
                                checked={selectedRoles.includes(role)}
                                onChange={() => handleRoleToggle(role)}
                                className="mr-2"
                            />
                            {role}
                        </label>
                    ))}
                </div>

                <div className="mb-6">
                    <h4 className="font-medium mb-2">Industry</h4>
                    {["SaaS", "Machine learning", "Crypto", "Retail technology"].map((industry) => (
                        <label key={industry} className="flex items-center mb-2">
                            <input
                                type="checkbox"
                                checked={selectedIndustries.includes(industry)}
                                onChange={() => handleIndustryToggle(industry)}
                                className="mr-2"
                            />
                            {industry}
                        </label>
                    ))}
                </div>
            </aside>

            {/* Job Collections Section */}
            <main className="flex-grow">
                <h2 className="text-2xl font-semibold mb-6">Job collections</h2>
                <div className="space-y-4">
                    {collections.map((collection, index) => (
                        <JobCard
                            key={index}
                            title={collection.title}
                            description={collection.description}
                            locations={collection.locations}
                            companies={collection.companies}
                            jobs={collection.jobs}
                            date={collection.date}
                        />
                    ))}
                </div>
            </main>
        </div>
    );
};

export default JobFeedComponent;
