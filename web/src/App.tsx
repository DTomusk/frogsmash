import { useItems } from "./hooks/useItems"

function App() {
  const { isPending, error, data } = useItems();

  if (isPending) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error loading items.</div>;
  }

  const { left_item, right_item } = data;

  return (
    <div className="flex gap-4">
      <img src={left_item.image_url} alt={left_item.name} />
      <img src={right_item.image_url} alt={right_item.name} />
    </div>
  );
}

export default App;
