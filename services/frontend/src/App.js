import React, {Component} from 'react';
import "./App.css";
import { ApolloProvider, useMutation } from "@apollo/react-hooks";
import { client } from "./apollo-client";
import { Route } from "react-router";
import {
    GetAllRestaurants,
    GetRestaurantsByID
} from "./gqlComponents/restaurant/RestaurantContainer";
import {
    GetAllCategories,
    GetCategoryByID
} from "./gqlComponents/category/CategoryContainer";
import { useState, useEffect } from "react";

const CategoryPage = ({match}) => {
    const {
        params: { categoryId },
    } = match;
    const [isLoading, setIsLoading] = useState(true);
    const [data, setData] = useState();

    useEffect(() => {
    },[categoryId]);
    return (
        <div>
            <ApolloProvider client={client}>
                <GetCategoryByID filter={categoryId} />
            </ApolloProvider>
        </div>
    )
}

const RestaurantPage = ({match}) => {
  const {
    params: { restaurantId },
  } = match;
  const [isLoading, setIsLoading] = useState(true);
  const [data, setData] = useState();

  useEffect(() => {
  },[restaurantId]);
  return (
    <div>
      <ApolloProvider client={client}>
        <GetRestaurantsByID filter={restaurantId} />
      </ApolloProvider>
    </div>
  )
}

function Categories() {
    return (
        <div>
            <ApolloProvider client={client}>
                <GetAllCategories />
            </ApolloProvider>
        </div>
    );
}

function Restaurants() {
  return (
    <div>
      <ApolloProvider client={client}>
        <GetAllRestaurants />
      </ApolloProvider>
    </div>
  );
}


class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let button;
    button = (
      <div>
        <Route exact path="/restaurant" component={Restaurants}>
          <Restaurants />
        </Route>
        <Route path="/restaurant/:restaurantId" component={RestaurantPage} />

        <Route exact path="/category" component={Categories}>
          <Categories />
        </Route>
        <Route path="/category/:categoryId" component={CategoryPage} />
      </div>
    );

    return (
      <div>
        {button}
      </div>
    );
  }
}

export default App;
