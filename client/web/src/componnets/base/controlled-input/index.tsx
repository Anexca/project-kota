import { Controller, ControllerProps, FieldValues } from "react-hook-form";
import { Input, InputProps } from "../input/input";

interface ControlledInputProps<TFieldValues extends FieldValues>
  extends Omit<ControllerProps<TFieldValues>, "render"> {
  inputProps?: InputProps;
}

const ControlledInput = <TFieldValues extends FieldValues>(
  props: ControlledInputProps<TFieldValues>
) => {
  const { inputProps, ...rest } = props;
  return (
    <Controller
      {...rest}
      render={({ field, fieldState }) => {
        return (
          <div>
            <Input {...field} {...inputProps} />
            {fieldState.error?.message ? (
              <div className="text-sm text-destructive">
                {fieldState.error?.message}
              </div>
            ) : null}
          </div>
        );
      }}
    />
  );
};

export default ControlledInput;
