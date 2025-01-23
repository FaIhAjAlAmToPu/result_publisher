export default function ResetPassword() {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <div className="bg-white p-6 rounded shadow-md w-full max-w-md">
          <h1 className="text-2xl font-bold mb-4 text-center">Reset Password</h1>
          <p className="text-sm text-gray-600 mb-4 text-center">
            Enter a new password for your account.
          </p>
          <form>
            <div className="mb-4">
              <label htmlFor="password" className="block text-sm font-medium text-gray-700">
                New Password
              </label>
              <input
                type="password"
                id="password"
                className="mt-1 block w-full p-2 border border-gray-300 rounded"
                placeholder="Enter your new password"
              />
            </div>
            <div className="mb-4">
              <label htmlFor="confirm-password" className="block text-sm font-medium text-gray-700">
                Confirm Password
              </label>
              <input
                type="password"
                id="confirm-password"
                className="mt-1 block w-full p-2 border border-gray-300 rounded"
                placeholder="Confirm your new password"
              />
            </div>
            <button
              type="submit"
              className="w-full bg-green-600 text-white py-2 px-4 rounded hover:bg-green-700"
            >
              Reset Password
            </button>
          </form>
          <p className="text-sm text-center mt-4">
            Remember your password?{" "}
            <a href="/login" className="text-blue-600 hover:underline">
              Login
            </a>
          </p>
        </div>
      </div>
    );
  }
  