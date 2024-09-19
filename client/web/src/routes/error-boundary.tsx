import { Component, PropsWithChildren } from "react";

interface State {
  hasError: boolean;
  errorMessage: string;
}
class ErrorBoundary extends Component<PropsWithChildren<any>, State> {
  constructor(props: any) {
    super(props);
    this.state = { hasError: false, errorMessage: "" };
  }

  // This method is invoked if any child component throws an error
  static getDerivedStateFromError(_: Error) {
    // Update state so the next render will show the fallback UI
    return { hasError: true };
  }

  // You can log the error details or perform any side effects here
  //   componentDidCatch(error: Error, errorInfo: ErrorInfo) {
  //     // this.setState({  errorMessage: error.toString() });
  //   }

  render() {
    if (this.state.hasError) {
      // Render fallback UI when an error occurs
      return <ErrorPage />;
    }

    // If no error, render children components as usual
    return this.props.children;
  }
}

export default ErrorBoundary;

function ErrorPage() {
  return (
    <>
      <main className="flex items-center justify-center min-h-screen   bg-white ">
        <div className="text-center">
          <h1 className="mt-4 text-2xl font-bold tracking-tight text-gray-900 sm:text-5xl">
            Oops an Error Occurred
          </h1>
          <p className="mt-6 text-base leading-7 text-gray-600">
            Sorry, there is some error on our side.
          </p>
          <div className="mt-10 flex items-center justify-center gap-x-6">
            <a
              href="/"
              className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              Go back home
            </a>
          </div>
        </div>
      </main>
    </>
  );
}
