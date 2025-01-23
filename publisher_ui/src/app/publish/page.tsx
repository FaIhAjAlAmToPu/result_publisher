'use client';
import { useState } from "react";

export default function UploadCSV() {
  const [examName, setExamName] = useState("");
  const [examFormat, setExamFormat] = useState("");
  const [publishDate, setPublishDate] = useState("");
  const [endTime, setEndTime] = useState("");
  const [csvFile, setCsvFile] = useState<File | null>(null);
  const [fileSizeError, setFileSizeError] = useState("");
  const [dragging, setDragging] = useState(false);

  const maxFileSize = 50 * 1024 * 1024; // 50MB max file size

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files ? e.target.files[0] : null;
    if (file) {
      if (file.size > maxFileSize) {
        setFileSizeError("File size exceeds the 50MB limit. Please upload a smaller file.");
        setCsvFile(null);
      } else {
        setFileSizeError("");
        setCsvFile(file);
      }
    }
  };

  const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();

    const file = e.dataTransfer.files ? e.dataTransfer.files[0] : null;
    if (file) {
      if (file.size > maxFileSize) {
        setFileSizeError("File size exceeds the 50MB limit. Please upload a smaller file.");
        setCsvFile(null);
      } else {
        setFileSizeError("");
        setCsvFile(file);
      }
    }

    setDragging(false);
  };

  const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setDragging(true);
  };

  const handleDragLeave = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    setDragging(false);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (!csvFile) {
      alert("Please upload a CSV file.");
      return;
    }

    // Handle file upload and form data processing here
    const formData = new FormData();
    formData.append("examName", examName);
    formData.append("examFormat", examFormat);
    formData.append("publishDate", publishDate);
    formData.append("endTime", endTime);
    formData.append("csvFile", csvFile);

    // You can make an API call here to submit the form data, for example:
    // fetch('/api/upload-results', { method: 'POST', body: formData });

    console.log("Form submitted with data:", formData);
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-6 rounded shadow-md w-full max-w-lg">
        <h1 className="text-3xl font-bold mb-6 text-center">Upload Exam Results</h1>
        
        <form onSubmit={handleSubmit}>
          {/* Exam Name */}
          <div className="mb-4">
            <label htmlFor="examName" className="block text-sm font-medium text-gray-700">
              Exam Name
            </label>
            <input
              type="text"
              id="examName"
              className="mt-1 block w-full p-3 border border-gray-300 rounded"
              placeholder="Enter Exam Name"
              value={examName}
              onChange={(e) => setExamName(e.target.value)}
              required
            />
          </div>

          {/* Exam Format */}
          <div className="mb-4">
            <label htmlFor="examFormat" className="block text-sm font-medium text-gray-700">
              Exam Format
            </label>
            <textarea
              id="examFormat"
              className="mt-1 block w-full p-3 border border-gray-300 rounded"
              placeholder="Describe the exam format, subjects, and marks"
              value={examFormat}
              onChange={(e) => setExamFormat(e.target.value)}
              required
            />
          </div>

          {/* Publish Date */}
          <div className="mb-4">
            <label htmlFor="publishDate" className="block text-sm font-medium text-gray-700">
              Publish Date
            </label>
            <input
              type="date"
              id="publishDate"
              className="mt-1 block w-full p-3 border border-gray-300 rounded"
              value={publishDate}
              onChange={(e) => setPublishDate(e.target.value)}
              required
            />
          </div>

          {/* End Time */}
          <div className="mb-4">
            <label htmlFor="endTime" className="block text-sm font-medium text-gray-700">
              End Time
            </label>
            <input
              type="time"
              id="endTime"
              className="mt-1 block w-full p-3 border border-gray-300 rounded"
              value={endTime}
              onChange={(e) => setEndTime(e.target.value)}
              required
            />
          </div>

          {/* CSV File Upload (Drag and Drop or Choose File) */}
          <div
            className={`mb-4 p-4 border-2 border-dashed rounded ${dragging ? "bg-gray-200" : ""}`}
            onDrop={handleDrop}
            onDragOver={handleDragOver}
            onDragLeave={handleDragLeave}
          >
            <label htmlFor="csvFile" className="block text-sm font-medium text-gray-700 mb-2">
              Upload CSV File
            </label>
            <input
              type="file"
              id="csvFile"
              accept=".csv"
              className="block w-full p-3 border border-gray-300 rounded mb-2"
              onChange={handleFileChange}
            />
            <p className="text-gray-500 text-sm">
              Drag and drop your CSV file here or click to choose a file.
            </p>
            {fileSizeError && <p className="text-red-500 text-sm mt-2">{fileSizeError}</p>}
          </div>

          {/* Submit Button */}
          <button
            type="submit"
            className="w-full bg-blue-600 text-white py-3 px-4 rounded hover:bg-blue-700"
          >
            Upload Results
          </button>
        </form>
      </div>
    </div>
  );
}
