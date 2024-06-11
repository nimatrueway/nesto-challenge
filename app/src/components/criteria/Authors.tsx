import React, {useState, useEffect, useCallback} from "react";
import Autocomplete from "@material-ui/lab/Autocomplete";
import TextField from "@material-ui/core/TextField";
import * as models from "../../models";
import {debounce} from 'lodash';

export default function Authors(props: {
  onChange: (authors: models.Author[]) => void;
}) {
  const FETCH_DEBOUNCE_DELAY = 250;

  const [error, setError] = React.useState<models.Error | null>(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [authors, setAuthors] = React.useState<models.Author[]>([]);
  const [inputValue, setInputValue] = useState('');

  const updateAuthors = useCallback(debounce((search: string) =>
          fetch(`http://localhost:5001/api/v1/authors?search=${search}`)
              .then(res => res.json())
              .then(
                  result => {
                    setIsLoaded(true);
                    setAuthors(result);
                  },
                  error => {
                    setIsLoaded(true);
                    setError(error);
                  }
              )
      , FETCH_DEBOUNCE_DELAY, {trailing: true}), [])

  useEffect(() => {
    updateAuthors(inputValue);
  }, [inputValue]);

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  if (!isLoaded) {
    return <div>Loading...</div>;
  }

  return (
      <Autocomplete
          multiple
          options={authors}
          getOptionLabel={x => `${x.firstName} ${x.lastName}`}
          filterOptions={(options, _) => options}
          onChange={(_, items) => props.onChange(items)}
          renderInput={params => (
              <TextField
                  {...params}
                  variant="standard"
                  label="Authors"
                  placeholder="Type or select author(s)"
                  onChange={(event) => setInputValue(event.target.value)}
              />
          )}
      />
  );
}
