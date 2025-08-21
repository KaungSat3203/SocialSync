"use client";

import { useEffect, useState } from "react";
import axios from "axios";

export default function FacebookAnalytics({ postID }) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!postID) return;

    axios
      .get(`/api/facebook/analytics?post_id=${postID}`)
      .then((res) => setData(res.data))
      .catch((err) => {
        console.error("Error fetching analytics:", err);
        setData(null);
      })
      .finally(() => setLoading(false));
  }, [postID]);

  if (loading) return <p>Loading analytics...</p>;
  if (!data) return <p>Analytics not available.</p>;

  return (
    <div className="p-4 bg-white rounded-xl shadow-md">
      <h2 className="text-xl font-bold mb-3">Facebook Post Analytics</h2>
      <ul className="space-y-1">
        {Object.entries(data).map(([key, value]) => (
          <li key={key} className="text-sm">
            <strong>{key}:</strong>{" "}
            {typeof value === "object" ? JSON.stringify(value) : value}
          </li>
        ))}
      </ul>
    </div>
  );
}
