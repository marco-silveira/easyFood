import React from "react";

const Categories = ({ loading, error, data }) => {
    if (loading) return "Loading...";
    if (error) return `Error! ${error.message}`;
    return (
        <ul>
            {data.category.map(category => (
                <li onClick={ () => {window.location.pathname === "/category/" ?  window.location =  category.id : window.location = "category/"+category.id}
                }>{category.name}</li>
            ))}
        </ul>
    );
};

const Category = ({ loading, error, data }) => {
    if (loading) return "Loading...";
    if (error) return `Error! ${error.message}`;
    return (
        <>
            <h1>{data.category[0].name}</h1>
            <h2>
                {data.category[0].restaurants && data.category[0].restaurants.map(restaurant => (
                    <li onClick={ () => {window.location = restaurant.id}}>{restaurant.name} - {restaurant.description}</li>
                ))}
            </h2>
        </>
    );
};

export { Categories, Category };
