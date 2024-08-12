import React from "react";

const SelectPct = () => {
  const [data, setData] = React.useState<{} | any>(null);

  const URL = "http://127.0.0.1:8010/microservice/titles";

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(URL, { method: "GET" });
        const result = await response.json();
        setData(result.body);
      } catch (err) {
        console.log(err);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="bg-black text-white">
      {data ? (
        <ul>
          {data.map((item, index) => (
            <li key={index}>
              <strong>Name:</strong> {item.Name} - <strong>Length:</strong>{" "}
              {item.Len}
            </li>
          ))}
        </ul>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default SelectPct;
