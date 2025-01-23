export default function SearchResults() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-6 rounded shadow-md w-full max-w-lg">
        <h1 className="text-3xl font-bold mb-6 text-center">Search Your Results</h1>
        
        <form>
          {/* Student ID Option */}
          <div className="mb-4">
            <label htmlFor="studentId" className="block text-sm font-medium text-gray-700">
              Student ID
            </label>
            <input
              type="text"
              id="studentId"
              className="mt-1 block w-full p-3 border border-gray-300 rounded"
              placeholder="Enter your Student ID"
            />
          </div>

          {/* Group ID Option */}
          <div className="mb-4">
            <label htmlFor="groupId" className="block text-sm font-medium text-gray-700">
              Group ID
            </label>
            <input
              type="text"
              id="groupId"
              className="mt-1 block w-full p-3 border border-gray-300 rounded"
              placeholder="Enter your Group ID"
            />
          </div>

          {/* Search Button */}
          <button 
            type="submit"
            className="w-full bg-green-600 text-white py-3 px-4 rounded hover:bg-green-700"
          >
            Search Results
          </button>
        </form>
      </div>
    </div>
  );
}
