import FacebookAnalytics from "../../components/FacebookAnalytics";

export default function PostDetailsPage() {
  const postID = "1234567890"; // Replace with actual post ID from backend

  return (
    <div className="max-w-xl mx-auto mt-10">
      <FacebookAnalytics postID={postID} />
    </div>
  );
}
