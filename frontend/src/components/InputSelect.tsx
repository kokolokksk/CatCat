import { useRef } from 'react';

const InputSelect = (prop: any) => {
  const { label, value, options } = prop;
  const inputRef = useRef<HTMLInputElement>(null);
  const selectRef = useRef<HTMLSelectElement>(null);

  const onChange = () => {
    inputRef.current!.value = selectRef.current?.value || '';
  };
  const addAndSet = () => {
    if (inputRef.current?.value) {
      const option = document.createElement('option');
      option.value = inputRef.current.value;
      option.text = inputRef.current.value;
      selectRef.current?.add(option);
      if (selectRef.current) {
        selectRef.current.value = inputRef.current.value;
      }
    }
  };
  const deleteOption = () => {
    if (selectRef.current) {
      selectRef.current.remove(selectRef.current.selectedIndex);
    }
  };
  return (
    <div className="input-select">
      <input type="text" ref={inputRef} />
      <button type="button" onClick={addAndSet}>
        确定
      </button>
      <p />
      <label htmlFor="select">{label}</label>
      <select name="select" value={value} onChange={onChange} ref={selectRef}>
        {options.map((option: string) => (
          <>
            <option value={option}>
              {option}
              <button type="button" onClick={deleteOption}>
                x
              </button>
            </option>
          </>
        ))}
      </select>
    </div>
  );
};

export default InputSelect;
