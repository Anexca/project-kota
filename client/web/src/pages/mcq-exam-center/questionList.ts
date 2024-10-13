export const questionsList = [
  {
    question:
      "A ball is thrown vertically upwards with an initial velocity (v0). Assuming no air resistance, what is the maximum height (h) reached by the ball?",
    options: [
      { label: "\\( \\frac{v_0^2}{2g} \\)", value: "v0^2/2g" },
      { label: "\\( \\frac{v_0^2}{g} \\)", value: "v0^2/g" },
      { label: "\\( \\frac{v_0}{2g} \\)", value: "v0/2g" },
      { label: "\\( \\frac{v_0}{g} \\)", value: "v0/g" },
    ],
  },
  {
    question:
      "What is the final velocity of a falling object after falling for time t, starting from rest? (Neglect air resistance)",
    options: [
      { label: "\\( gt \\)", value: "gt" },
      { label: "\\( \\frac{1}{2}gt^2 \\)", value: "1/2 gt^2" },
      { label: "\\( \\frac{1}{2}v_0t \\)", value: "1/2 v0 t" },
      { label: "\\( v_0 + gt \\)", value: "v0 + gt" },
    ],
  },
  {
    question:
      "A car accelerates from rest at a constant rate a for time t. What is the final velocity of the car?",
    options: [
      { label: "\\( at \\)", value: "at" },
      { label: "\\( \\frac{1}{2}at^2 \\)", value: "1/2 at^2" },
      { label: "\\( v_0 + at \\)", value: "v0 + at" },
      { label: "\\( v_0t + \\frac{1}{2}at^2 \\)", value: "v0 t + 1/2 at^2" },
    ],
  },
  {
    question:
      "What is the distance traveled by an object under constant acceleration a in time t, starting from rest?",
    options: [
      { label: "\\( \\frac{1}{2}at^2 \\)", value: "1/2 at^2" },
      { label: "\\( at \\)", value: "at" },
      { label: "\\( v_0t \\)", value: "v0 t" },
      { label: "\\( \\frac{1}{2}v_0t \\)", value: "1/2 v0 t" },
    ],
  },
  {
    question:
      "An object is thrown horizontally from a height h with initial speed v0. What is the time t it takes to hit the ground?",
    options: [
      { label: "\\( \\sqrt{\\frac{2h}{g}} \\)", value: "sqrt(2h/g)" },
      { label: "\\( \\frac{h}{v_0} \\)", value: "h/v0" },
      { label: "\\( \\frac{v_0}{g} \\)", value: "v0/g" },
      { label: "\\( \\frac{h}{2v_0} \\)", value: "h/2v0" },
    ],
  },
  {
    question:
      "What is the horizontal distance traveled by a projectile in time t with initial speed v0 at an angle \\( \\theta \\)?",
    options: [
      { label: "\\( v_0t \\cos \\theta \\)", value: "v0 t cos(theta)" },
      { label: "\\( v_0t \\sin \\theta \\)", value: "v0 t sin(theta)" },
      {
        label: "\\( \\frac{v_0^2 \\sin 2\\theta}{g} \\)",
        value: "v0^2 sin(2theta)/g",
      },
      { label: "\\( v_0^2 \\cos \\theta \\)", value: "v0^2 cos(theta)" },
    ],
  },
  {
    question:
      "For a projectile motion, what is the maximum height \\( h \\) reached by the projectile?",
    options: [
      {
        label: "\\( \\frac{v_0^2 \\sin^2 \\theta}{2g} \\)",
        value: "v0^2 sin^2(theta)/2g",
      },
      {
        label: "\\( \\frac{v_0^2 \\cos^2 \\theta}{g} \\)",
        value: "v0^2 cos^2(theta)/g",
      },
      { label: "\\( \\frac{v_0^2}{2g} \\)", value: "v0^2/2g" },
      { label: "\\( v_0 \\sin \\theta \\)", value: "v0 sin(theta)" },
    ],
  },
  {
    question:
      "What is the time of flight \\( T \\) for a projectile launched at an angle \\( \\theta \\) with initial speed \\( v_0 \\)?",
    options: [
      {
        label: "\\( \\frac{2v_0 \\sin \\theta}{g} \\)",
        value: "2v0 sin(theta)/g",
      },
      {
        label: "\\( \\frac{v_0 \\cos \\theta}{g} \\)",
        value: "v0 cos(theta)/g",
      },
      {
        label: "\\( \\frac{v_0^2 \\sin^2 \\theta}{g} \\)",
        value: "v0^2 sin^2(theta)/g",
      },
      {
        label: "\\( \\frac{v_0 \\sin \\theta}{g} \\)",
        value: "v0 sin(theta)/g",
      },
    ],
  },
  {
    question:
      "What is the final velocity of an object dropped from rest after falling for time t in a vacuum?",
    options: [
      { label: "\\( gt \\)", value: "gt" },
      { label: "\\( \\frac{1}{2}gt^2 \\)", value: "1/2 gt^2" },
      { label: "\\( v_0 + gt \\)", value: "v0 + gt" },
      { label: "\\( \\frac{1}{2}v_0t \\)", value: "1/2 v0 t" },
    ],
  },
  {
    question:
      "What is the distance traveled by an object in free fall in time t?",
    options: [
      { label: "\\( \\frac{1}{2}gt^2 \\)", value: "1/2 gt^2" },
      { label: "\\( gt \\)", value: "gt" },
      { label: "\\( v_0t \\)", value: "v0 t" },
      { label: "\\( \\frac{1}{2}v_0t \\)", value: "1/2 v0 t" },
    ],
  },
  {
    question:
      "A projectile is launched with initial speed \\( v_0 \\) at an angle \\( \\theta \\). What is the range R of the projectile?",
    options: [
      {
        label: "\\( \\frac{v_0^2 \\sin 2\\theta}{g} \\)",
        value: "v0^2 sin(2theta)/g",
      },
      { label: "\\( v_0^2 \\cos \\theta \\)", value: "v0^2 cos(theta)" },
      {
        label: "\\( \\frac{v_0^2 \\cos^2 \\theta}{g} \\)",
        value: "v0^2 cos^2(theta)/g",
      },
      {
        label: "\\( \\frac{v_0^2 \\sin^2 \\theta}{2g} \\)",
        value: "v0^2 sin^2(theta)/2g",
      },
    ],
  },
  {
    question:
      "If an object is projected upward with initial speed v0, what is the time to reach maximum height?",
    options: [
      { label: "\\( \\frac{v_0}{g} \\)", value: "v0/g" },
      { label: "\\( \\frac{v_0^2}{2g} \\)", value: "v0^2/2g" },
      { label: "\\( \\frac{2v_0}{g} \\)", value: "2v0/g" },
      { label: "\\( \\frac{v_0}{2g} \\)", value: "v0/2g" },
    ],
  },
  {
    question:
      "An object moves with constant velocity v0. What is the displacement after time t?",
    options: [
      { label: "\\( v_0t \\)", value: "v0 t" },
      { label: "\\( \\frac{1}{2}v_0t^2 \\)", value: "1/2 v0 t^2" },
      { label: "\\( v_0^2t \\)", value: "v0^2 t" },
      { label: "\\( \\frac{v_0}{t} \\)", value: "v0 / t" },
    ],
  },
  {
    question:
      "If an object is dropped from height h, what is its velocity just before hitting the ground?",
    options: [
      { label: "\\( \\sqrt{2gh} \\)", value: "sqrt(2gh)" },
      { label: "\\( gh \\)", value: "gh" },
      { label: "\\( \\frac{h}{2g} \\)", value: "h/2g" },
      { label: "\\( \\frac{v_0}{g} \\)", value: "v0/g" },
    ],
  },
  {
    question:
      "A body moves in a straight line with constant acceleration a. If its initial velocity is v0, what is its velocity after time t?",
    options: [
      { label: "\\( v_0 + at \\)", value: "v0 + at" },
      { label: "\\( v_0t + \\frac{1}{2}at^2 \\)", value: "v0 t + 1/2 at^2" },
      { label: "\\( \\frac{v_0^2}{a} \\)", value: "v0^2/a" },
      { label: "\\( v_0 \\)", value: "v0" },
    ],
  },
  {
    question:
      "What is the acceleration of an object moving with constant velocity?",
    options: [
      { label: "\\( 0 \\)", value: "0" },
      { label: "\\( v_0 \\)", value: "v0" },
      { label: "\\( g \\)", value: "g" },
      { label: "\\( \\frac{v_0}{t} \\)", value: "v0/t" },
    ],
  },
  {
    question:
      "A car accelerates from rest at a constant rate a. What is the distance traveled by the car after time t?",
    options: [
      { label: "\\( \\frac{1}{2}at^2 \\)", value: "1/2 at^2" },
      { label: "\\( at^2 \\)", value: "at^2" },
      { label: "\\( v_0t \\)", value: "v0 t" },
      { label: "\\( \\frac{1}{2}v_0t \\)", value: "1/2 v0 t" },
    ],
  },
  {
    question:
      "A stone is thrown upwards with initial velocity v0. What is the total time of flight until it returns to the original position?",
    options: [
      { label: "\\( \\frac{2v_0}{g} \\)", value: "2v0/g" },
      { label: "\\( \\frac{v_0}{g} \\)", value: "v0/g" },
      { label: "\\( v_0 \\)", value: "v0" },
      { label: "\\( \\frac{1}{2}v_0 \\)", value: "1/2 v0" },
    ],
  },
  {
    question:
      "For an object projected horizontally with speed v0 from height h, what is its range R?",
    options: [
      { label: "\\( v_0 \\sqrt{\\frac{2h}{g}} \\)", value: "v0 sqrt(2h/g)" },
      {
        label: "\\( \\frac{v_0^2 \\sin \\theta}{g} \\)",
        value: "v0^2 sin(theta)/g",
      },
      { label: "\\( v_0 \\frac{h}{g} \\)", value: "v0 h/g" },
      { label: "\\( v_0t \\)", value: "v0 t" },
    ],
  },
  {
    question:
      "What is the centripetal acceleration of an object moving in a circle of radius r with speed v?",
    options: [
      { label: "\\( \\frac{v^2}{r} \\)", value: "v^2/r" },
      { label: "\\( vr \\)", value: "vr" },
      { label: "\\( \\frac{v}{r} \\)", value: "v/r" },
      { label: "\\( \\frac{v^2 r}{g} \\)", value: "v^2 r/g" },
    ],
  },
  {
    question:
      "What is the force acting on an object of mass m moving with acceleration a?",
    options: [
      { label: "\\( ma \\)", value: "ma" },
      { label: "\\( m + a \\)", value: "m + a" },
      { label: "\\( m - a \\)", value: "m - a" },
      { label: "\\( ma^2 \\)", value: "ma^2" },
    ],
  },
  {
    question:
      "If an object of mass m is moving with initial velocity v0 and it experiences a constant force F, what is the final velocity after time t?",
    options: [
      { label: "\\( v_0 + \\frac{F}{m} t \\)", value: "v0 + (F/m) t" },
      { label: "\\( v_0 - \\frac{F}{m} t \\)", value: "v0 - (F/m) t" },
      { label: "\\( v_0 \\frac{F}{m} t \\)", value: "v0 (F/m) t" },
      { label: "\\( v_0 + Ft \\)", value: "v0 + Ft" },
    ],
  },
  {
    question:
      "A body is thrown upward with an initial velocity v0. What is the velocity of the body at height h above the ground?",
    options: [
      { label: "\\( \\sqrt{v_0^2 - 2gh} \\)", value: "sqrt(v0^2 - 2gh)" },
      { label: "\\( v_0 - gh \\)", value: "v0 - gh" },
      { label: "\\( v_0 \\)", value: "v0" },
      { label: "\\( \\frac{v_0}{2} \\)", value: "v0/2" },
    ],
  },
  {
    question:
      "What is the work done by a force F in moving an object a distance d?",
    options: [
      { label: "\\( Fd \\)", value: "Fd" },
      { label: "\\( F + d \\)", value: "F + d" },
      { label: "\\( F - d \\)", value: "F - d" },
      { label: "\\( \\frac{F}{d} \\)", value: "F/d" },
    ],
  },
  {
    question:
      "An object of mass m is dropped from a height h. What is the potential energy of the object at height h?",
    options: [
      { label: "\\( mgh \\)", value: "mgh" },
      { label: "\\( mgh^2 \\)", value: "mgh^2" },
      { label: "\\( \\frac{1}{2} mgh \\)", value: "1/2 mgh" },
      { label: "\\( mg \\)", value: "mg" },
    ],
  },
  {
    question:
      "For an object of mass m undergoing uniform circular motion with radius r and speed v, what is the centripetal force?",
    options: [
      { label: "\\( \\frac{mv^2}{r} \\)", value: "mv^2/r" },
      { label: "\\( mvr \\)", value: "mvr" },
      { label: "\\( m \\frac{v}{r} \\)", value: "m v/r" },
      { label: "\\( \\frac{v^2}{mr} \\)", value: "v^2/mr" },
    ],
  },
  {
    question:
      "What is the total energy of an object in free fall just before hitting the ground?",
    options: [
      { label: "\\( mgh \\)", value: "mgh" },
      { label: "\\( \\frac{1}{2}mv^2 \\)", value: "1/2 mv^2" },
      { label: "\\( mgh + \\frac{1}{2}mv^2 \\)", value: "mgh + 1/2 mv^2" },
      { label: "\\( \\frac{1}{2}mgh \\)", value: "1/2 mgh" },
    ],
  },
  {
    question:
      "An object with mass m is moving with velocity v0. What is its kinetic energy?",
    options: [
      { label: "\\( \\frac{1}{2}mv_0^2 \\)", value: "1/2 mv0^2" },
      { label: "\\( mv_0^2 \\)", value: "mv0^2" },
      { label: "\\( mgh \\)", value: "mgh" },
      { label: "\\( \\frac{1}{2}mv_0 \\)", value: "1/2 mv0" },
    ],
  },
  {
    question:
      "What is the final kinetic energy of an object after it has been accelerated by a constant force F over distance d starting from rest?",
    options: [
      { label: "\\( Fd \\)", value: "Fd" },
      { label: "\\( \\frac{Fd^2}{2m} \\)", value: "Fd^2/2m" },
      { label: "\\( \\frac{F^2 d}{2m} \\)", value: "F^2 d/2m" },
      { label: "\\( \\frac{1}{2} Fd \\)", value: "1/2 Fd" },
    ],
  },
  {
    question:
      "If a force F is applied to an object of mass m and causes an acceleration a, what is the work done if the object moves a distance d?",
    options: [
      { label: "\\( Fd \\)", value: "Fd" },
      { label: "\\( ma \\)", value: "ma" },
      { label: "\\( \\frac{1}{2} Fd \\)", value: "1/2 Fd" },
      { label: "\\( Fd - ma \\)", value: "Fd - ma" },
    ],
  },
  {
    question:
      "What is the velocity of an object after falling freely from height h with initial velocity 0?",
    options: [
      { label: "\\( \\sqrt{2gh} \\)", value: "sqrt(2gh)" },
      { label: "\\( gh \\)", value: "gh" },
      { label: "\\( \\frac{h}{2g} \\)", value: "h/2g" },
      { label: "\\( \\frac{v_0}{g} \\)", value: "v0/g" },
    ],
  },
  {
    question:
      "For an object of mass m falling under gravity with acceleration g, what is the force of gravity acting on it?",
    options: [
      { label: "\\( mg \\)", value: "mg" },
      { label: "\\( m + g \\)", value: "m + g" },
      { label: "\\( m - g \\)", value: "m - g" },
      { label: "\\( m/g \\)", value: "m/g" },
    ],
  },
  {
    question:
      "A projectile is launched with speed v0 at an angle \\( \\theta \\). What is the vertical component of the velocity?",
    options: [
      { label: "\\( v_0 \\sin \\theta \\)", value: "v0 sin(theta)" },
      { label: "\\( v_0 \\cos \\theta \\)", value: "v0 cos(theta)" },
      { label: "\\( v_0^2 \\sin^2 \\theta \\)", value: "v0^2 sin^2(theta)" },
      { label: "\\( v_0 \\)", value: "v0" },
    ],
  },
  {
    question:
      "A mass m is hung from a spring and extends it by distance x. What is the force exerted by the spring?",
    options: [
      { label: "\\( kx \\)", value: "kx" },
      { label: "\\( mg \\)", value: "mg" },
      { label: "\\( m + x \\)", value: "m + x" },
      { label: "\\( \frac{k}{x} \\)", value: "k/x" },
    ],
  },
  {
    question:
      "What is the change in potential energy of an object of mass m when it is lifted to height h?",
    options: [
      { label: "\\( mgh \\)", value: "mgh" },
      { label: "\\( mg - h \\)", value: "mg - h" },
      { label: "\\( mg \\)", value: "mg" },
      { label: "\\( \\frac{1}{2} mgh \\)", value: "1/2 mgh" },
    ],
  },
  {
    question:
      "For a body moving in a circle of radius r with constant speed v, what is the centripetal force?",
    options: [
      { label: "\\( \\frac{mv^2}{r} \\)", value: "mv^2/r" },
      { label: "\\( mv \\)", value: "mv" },
      { label: "\\( \\frac{v}{r} \\)", value: "v/r" },
      { label: "\\( \\frac{v^2}{r} \\)", value: "v^2/r" },
    ],
  },
  {
    question: "The work done by a force F over distance d is?",
    options: [
      { label: "\\( Fd \\)", value: "Fd" },
      { label: "\\( F + d \\)", value: "F + d" },
      { label: "\\( F - d \\)", value: "F - d" },
      { label: "\\( F/d \\)", value: "F/d" },
    ],
  },
  {
    question:
      "What is the kinetic energy of an object with mass m moving at speed v?",
    options: [
      { label: "\\( \\frac{1}{2}mv^2 \\)", value: "1/2 mv^2" },
      { label: "\\( mv \\)", value: "mv" },
      { label: "\\( \\frac{1}{2}mv \\)", value: "1/2 mv" },
      { label: "\\( mv^2 \\)", value: "mv^2" },
    ],
  },
  {
    question:
      "If an object of mass m is dropped from rest, what is its velocity after falling for time t?",
    options: [
      { label: "\\( gt \\)", value: "gt" },
      { label: "\\( \\frac{1}{2}gt^2 \\)", value: "1/2 gt^2" },
      { label: "\\( v_0t \\)", value: "v0 t" },
      { label: "\\( v_0 + gt \\)", value: "v0 + gt" },
    ],
  },
  {
    question:
      "For a body in projectile motion, what is the range of the projectile?",
    options: [
      {
        label: "\\( \\frac{v_0^2 \\sin 2\\theta}{g} \\)",
        value: "v0^2 sin(2theta)/g",
      },
      {
        label: "\\( \\frac{v_0 \\sin \\theta}{g} \\)",
        value: "v0 sin(theta)/g",
      },
      {
        label: "\\( \\frac{v_0^2 \\cos^2 \\theta}{g} \\)",
        value: "v0^2 cos^2(theta)/g",
      },
      { label: "\\( v_0 t \\)", value: "v0 t" },
    ],
  },
];
