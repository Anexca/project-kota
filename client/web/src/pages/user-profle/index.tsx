import { yupResolver } from "@hookform/resolvers/yup";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { Link } from "react-router-dom";
import { Button } from "../../componnets/base/button/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "../../componnets/base/card/card";
import ControlledInput from "../../componnets/base/controlled-input";
import { Label } from "../../componnets/base/label/label";
import Loader from "../../componnets/shared/loder";
import { useToast } from "../../hooks/use-toast";
import { paths } from "../../routes/route.constant";
import { updateUserProfile } from "../../services/profile.service";
import useUserProfileStore from "../../store/user-info-store";
import {
  UserProfileSchema,
  UserProfileType,
} from "../../validation-schema/user-profile-schema";
import { Input } from "../../componnets/base/input/input";

const UserProfile = () => {
  const [loading, setIsloading] = useState(false);
  const { profile, setProfile } = useUserProfileStore();
  const { toast } = useToast();
  const { handleSubmit, control, reset } = useForm({
    defaultValues: {
      firstName: profile.first_name || "",
      lastName: profile.last_name || "",
      phoneNumber: profile.phone_number || "",
    },
    resolver: yupResolver(UserProfileSchema),
  });
  const onSumbit = async (formData: UserProfileType) => {
    setIsloading(true);
    try {
      const payload = {
        first_name: formData.firstName,
        last_name: formData.lastName,
        phone: formData.phoneNumber,
      };
      const res = await updateUserProfile(payload);
      const sanitizedRespomse = {
        firstName: res.data.first_name,
        lastName: res.data.last_name,
        phoneNumber: res.data.phone_number,
      };
      reset(sanitizedRespomse);
      setProfile(res.data);
      toast({
        title: "Successful",
        variant: "success",
        description: "Profile updated successfully.",
        duration: 5000,
      });
    } catch (error: any) {
      toast({
        title: error.message || "Something went wrong.",
        variant: "destructive",
        description: "Sorry there is some problem in proccessing your request.",
        duration: 5000,
      });
    } finally {
      setIsloading(false);
    }
  };
  return (
    <div className="flex items-center justify-center py-12">
      <Card className="max-w-[20rem]">
        <CardHeader>
          <CardTitle className="text-xl flex">
            Profile
            <Link className="ml-auto" to={`/${paths.HOMEPAGE}`}>
              <i className="fa-solid fa-xmark"></i>
            </Link>
          </CardTitle>
          <CardDescription>
            Complete your profile below to get subscription alerts and
            follow-up.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form
            onSubmit={handleSubmit(onSumbit)}
            noValidate
            className="grid gap-4"
          >
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="email">Email</Label>
              </div>
              <Input value={profile.email} disabled />
            </div>
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="email">First Name</Label>
              </div>
              <ControlledInput
                control={control}
                name="firstName"
                inputProps={{
                  placeholder: "Ramesh",
                }}
              />
            </div>
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="email">Last Name</Label>
              </div>
              <ControlledInput
                control={control}
                name="lastName"
                inputProps={{
                  placeholder: "Babu",
                }}
              />
            </div>
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="email">Phone</Label>
              </div>
              <ControlledInput
                control={control}
                name="phoneNumber"
                inputProps={{
                  placeholder: "8498776565",
                }}
              />
            </div>

            <Button
              type="submit"
              variant={"info"}
              disabled={loading}
              className="w-full"
            >
              {loading ? <Loader color={"outline"} /> : "Save"}
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
};

export default UserProfile;
