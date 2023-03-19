const people = [
  { name: 'Lindsay Walton', title: 'Front-end Developer', email: 'lindsay.walton@example.com', role: 'Member' },
  { name: 'Lindsay Walton', title: 'Front-end Developer', email: 'lindsay.walton@example.com', role: 'Member' },
  { name: 'Lindsay Walton', title: 'Front-end Developer', email: 'lindsay.walton@example.com', role: 'Member' },
  { name: 'Lindsay Walton', title: 'Front-end Developer', email: 'lindsay.walton@example.com', role: 'Member' },
]

export default function Table() {
  return (
    <div className="px-4 sm:px-6 lg:px-8">
      <div className="sm:flex sm:items-center">
        <div className="sm:flex-auto">
          <h1 className="text-base font-semibold leading-6 text-gray-900">Users</h1>
          <p className="mt-2 text-sm text-gray-700">
            A list of all the users in your account including their name, title, email and role.
          </p>
        </div>
        <div className="mt-4 sm:mt-0 sm:ml-16 sm:flex-nonex">
          <button
            type="button"
            className="block rounded-md bg-indigo-600 py-2 px-3 text-center text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Add user
          </button>
        </div>
      </div>
      <div className="mt-8 flow-root">
        <div className="overflow-x-auto ">
          <table>
            <thead className="">
              <tr className="divide-x divide-white bg-gray-200">
                <th scope="col" className="w-1/10 py-3.5 pl-4 pr-4 text-left text-sm font-semibold text-gray-900 sm:pl-0">
                  Player
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  1st
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  2nd
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  3rd
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  4th
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  Min
                </th>
              </tr>
            </thead>
            <tbody className="text-white bg-gray-200 divide-y divide-x divid-white">
              <tr>
                <td className="p-0">
                  <div className="h-full w-full bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
              </tr>
              <tr>
                <td className="p-0">
                  <div className="h-full w-full bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
                <td className="p-0">
                  <div className="h-full w-1/4 bg-blue-500 px-0 py-0">内容</div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div >
  )
}
