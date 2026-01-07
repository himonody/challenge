import withPWAInit from "next-pwa";

const withPWA = withPWAInit({
  dest: "public",
  disable: process.env.NODE_ENV === "development",
});

const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    const backend = process.env.BACKEND_API;
    if (!backend) {
      console.warn("BACKEND_API is not defined; skipping /proxy rewrites.");
      return [];
    }
    return [
      {
        source: "/proxy/:path*",
        destination: `${backend}/:path*`,
      },
    ];
  },
};

export default withPWA(nextConfig);
