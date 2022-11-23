export default async function HogeLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  throw new Promise((resolve) => setTimeout(resolve, 1000));
  return (
    <>
      <p>hoge layout</p>
      {children}
    </>
  );
}
